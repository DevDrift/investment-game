package models

import (
	"encoding/json"
	"github.com/DevDrift/investment-game/pkg/core"
	cached "github.com/DevDrift/investment-game/pkg/database"
	"github.com/google/uuid"
	"time"
)

const (
	AuctionTable          = "auctions"
	CompletedAuctionTable = "completed-auctions"
	AuctionTime           = time.Second * 30
)

type AuctionRequest struct {
	Auction *core.Auction
}

// New create new auction
func (auc *AuctionRequest) New(assetId string, step float64) (item *core.Auction, err error) {
	startTime := time.Now().UTC()
	item = &core.Auction{
		Id:        uuid.NewString(),
		StartTime: startTime,
		EndTime:   startTime.Add(AuctionTime),
		AssetId:   assetId,
		Step:      step,
	}
	return
}

// Add by id
func (auc *AuctionRequest) Add() (item *core.Auction, err error) {
	db, err := cached.OpenDb(AuctionTable)
	if err != nil {
		return
	}
	item = auc.Auction
	err = db.BitAdd(item.Key(), item.Value())
	return
}

// Get by id
func (auc *AuctionRequest) Get(id []byte) (item *core.Auction, err error) {
	db, err := cached.OpenDb(AuctionTable)
	if err != nil {
		return
	}
	err, bytes := db.BitGet(id)
	if err != nil {
		return
	}
	if bytes == nil {
		return
	}
	err = json.Unmarshal(bytes, &item)
	if err != nil {
		return
	}
	return
}

// List items
func (auc *AuctionRequest) List() (items []core.Auction, err error) {
	db, err := cached.OpenDb(AuctionTable)
	if err != nil {
		return
	}
	values, err := db.GetValues()
	if err != nil {
		return
	}
	for _, value := range values {
		var current core.Auction
		err = json.Unmarshal(value.Value, &current)
		if err != nil {
			continue
		}
		items = append(items, current)
	}
	return
}

// Update by id
func (auc *AuctionRequest) Update(id []byte) (item *core.Auction, err error) {
	db, err := cached.OpenDb(AuctionTable)
	if err != nil {
		return
	}
	item = auc.Auction
	err = db.BitAdd(id, item.Value())
	return
}

// Delete by id
func (auc *AuctionRequest) Delete(id []byte) (err error) {
	db, err := cached.OpenDb(AuctionTable)
	if err != nil {
		return
	}
	return db.Delete(id)
}

// Закртыть аукцион и определить победителя
