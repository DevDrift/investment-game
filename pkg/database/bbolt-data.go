package cached

import (
	"fmt"
	"github.com/DevDrift/investment-game/pkg/utils"
	bolt "go.etcd.io/bbolt"
	"sync"
)

var (
	mu          sync.Mutex
	DatabaseDir = utils.GetEnv("BACKUP_DIR", "ig-base")
	cacheData   = map[string]*Backup{}
)

type Item struct {
	Key   []byte
	Value []byte
}

type Backup struct {
	Bucket string
	*bolt.DB
}

func OpenDb(table string) (*Backup, error) {
	mu.Lock()
	defer mu.Unlock()
	basePath := fmt.Sprintf("%s/%s.igb", DatabaseDir, table)
	if cacheData[basePath] == nil {
		db, err := bolt.Open(basePath, 0600, &bolt.Options{Timeout: 0})
		if err != nil {
			return nil, err
		}
		cacheData[basePath] = &Backup{table, db}
		return cacheData[basePath], err
	}
	return cacheData[basePath], nil
}

func (db *Backup) BitAdd(key []byte, item []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(db.Bucket))
		if err != nil {
			return err
		}
		return bucket.Put(key, item)
	})
}

func (db *Backup) BitGet(key []byte) (error, []byte) {
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

func (db *Backup) GetValues() ([]Item, error) {
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

func (db *Backup) BitGetFist() (error, []byte) {
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

func (db *Backup) BitExists(key []byte) (error, bool) {
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
			return fmt.Errorf("key not found")
		}
		return nil
	})
	if err != nil && err.Error() != "key not found" {
		return err, false
	}
	if err != nil && err.Error() == "key not found" {
		return nil, false
	}
	return nil, true
}

func (db *Backup) Delete(key []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(db.Bucket))
		return bucket.Delete(key)
	})
}

func (db *Backup) DeleteBucket() error {
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

func (db *Backup) Close() error {
	return db.DB.Close()
}
