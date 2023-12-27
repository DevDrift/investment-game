package database

import (
	"fmt"
	"github.com/DevDrift/investment-game/pkg/utils"
	bolt "go.etcd.io/bbolt"
	"sync"
)

const (
	ErrKeyNotFound = "key not found"
)

var (
	mu        sync.Mutex
	Dir       = utils.GetEnv("DATABASE_DIR", "ig-base")
	cacheData = map[string]*Data{}
)

type Item struct {
	Key   []byte
	Value []byte
}

type Data struct {
	Bucket string
	*bolt.DB
}

func OpenDb(table string) (*Data, error) {
	mu.Lock()
	defer mu.Unlock()
	basePath := fmt.Sprintf("%s/%s.igb", Dir, table)
	if cacheData[basePath] == nil {
		db, err := bolt.Open(basePath, 0600, &bolt.Options{Timeout: 0})
		if err != nil {
			return nil, err
		}
		cacheData[basePath] = &Data{table, db}
		return cacheData[basePath], err
	}
	return cacheData[basePath], nil
}

func (db *Data) BitAdd(key []byte, item []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(db.Bucket))
		if err != nil {
			return err
		}
		return bucket.Put(key, item)
	})
}

func (db *Data) BucketAdd(bucket, key, item []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		bkt, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		return bkt.Put(key, item)
	})
}

func (db *Data) BitGet(key []byte) (error, []byte) {
	var result []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(db.Bucket))
		if bucket == nil {
			_, err := tx.CreateBucket([]byte(db.Bucket))
			if err != nil {
				return err
			}
		}
		bucket = tx.Bucket([]byte(db.Bucket))
		result = bucket.Get(key)
		return nil
	})
	if err != nil {
		return err, nil
	}
	return nil, result
}

func (db *Data) BucketGet(bucket, key []byte) (error, []byte) {
	var result []byte
	err := db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(bucket)
		if bkt == nil {
			_, err := tx.CreateBucket(bucket)
			if err != nil {
				return err
			}
		}
		bkt = tx.Bucket(bucket)
		result = bkt.Get(key)
		return nil
	})
	if err != nil {
		return err, nil
	}
	return nil, result
}

func (db *Data) GetValues() ([]Item, error) {
	totalScan := 0
	var result []Item
	_ = db.View(func(tx *bolt.Tx) error {
		totalScan++
		bucket := tx.Bucket([]byte(db.Bucket))
		if bucket == nil {
			_, err := tx.CreateBucket([]byte(db.Bucket))
			if err != nil {
				return err
			}
		}
		bucket = tx.Bucket([]byte(db.Bucket))
		bucket.ForEach(func(key, value []byte) error {
			result = append(result, Item{
				Key:   key,
				Value: value,
			})
			return nil
		})
		return nil
	})
	return result, nil
}

func (db *Data) BucketGetValues(bucket []byte) ([]Item, error) {
	totalScan := 0
	var result []Item
	_ = db.View(func(tx *bolt.Tx) error {
		totalScan++
		bkt := tx.Bucket(bucket)
		if bkt == nil {
			_, err := tx.CreateBucket(bucket)
			if err != nil {
				return err
			}
		}
		bkt = tx.Bucket(bucket)
		bkt.ForEach(func(key, value []byte) error {
			result = append(result, Item{
				Key:   key,
				Value: value,
			})
			return nil
		})
		return nil
	})
	return result, nil
}

func (db *Data) BitGetFist() (error, []byte) {
	var result []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(db.Bucket))
		if bucket == nil {
			_, err := tx.CreateBucket([]byte(db.Bucket))
			if err != nil {
				return err
			}
		}
		bucket = tx.Bucket([]byte(db.Bucket))
		cursor := bucket.Cursor()
		_, result = cursor.First()
		return nil
	})
	if err != nil {
		return err, nil
	}
	return nil, result
}

func (db *Data) BucketGetFist(bucket []byte) (error, []byte) {
	var result []byte
	err := db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(bucket)
		if bkt == nil {
			_, err := tx.CreateBucket(bucket)
			if err != nil {
				return err
			}
		}
		bkt = tx.Bucket(bucket)
		cursor := bkt.Cursor()
		_, result = cursor.First()
		return nil
	})
	if err != nil {
		return err, nil
	}
	return nil, result
}

func (db *Data) BitExists(key []byte) (error, bool) {
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(db.Bucket))
		if bucket == nil {
			_, err := tx.CreateBucket([]byte(db.Bucket))
			if err != nil {
				return err
			}
		}
		bucket = tx.Bucket([]byte(db.Bucket))
		if bucket.Get(key) == nil {
			return fmt.Errorf(ErrKeyNotFound)
		}
		return nil
	})
	if err != nil {
		if err.Error() != ErrKeyNotFound || err.Error() == ErrKeyNotFound {
			return err, false
		}
	}
	return nil, true
}

func (db *Data) BucketExists(bucket, key []byte) (error, bool) {
	err := db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(bucket)
		if bkt == nil {
			_, err := tx.CreateBucket(bucket)
			if err != nil {
				return err
			}
		}
		bkt = tx.Bucket(bucket)
		if bkt.Get(key) == nil {
			return fmt.Errorf(ErrKeyNotFound)
		}
		return nil
	})
	if err != nil {
		if err.Error() != ErrKeyNotFound || err.Error() == ErrKeyNotFound {
			return err, false
		}
	}
	return nil, true
}

func (db *Data) Delete(key []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(db.Bucket))
		return bucket.Delete(key)
	})
}

func (db *Data) BucketDeleteKey(bucket, key []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(bucket)
		return bkt.Delete(key)
	})
}

func (db *Data) DeleteBucket() error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(db.Bucket))
		if bucket == nil {
			_, err := tx.CreateBucket([]byte(db.Bucket))
			if err != nil {
				return err
			}
		}
		return tx.DeleteBucket([]byte(db.Bucket))
	})
}

func (db *Data) DropBucket(bucket []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(bucket)
		if bkt == nil {
			_, err := tx.CreateBucket(bucket)
			if err != nil {
				return err
			}
		}
		return tx.DeleteBucket(bucket)
	})
}

func (db *Data) Close() error {
	return db.DB.Close()
}
