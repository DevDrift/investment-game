package models

import (
	"encoding/json"
	"errors"
	"github.com/DevDrift/investment-game/pkg/core"
	cached "github.com/DevDrift/investment-game/pkg/database"
	"time"
)

type BidRequest struct {
	AuctionId string    `json:"auctionId"`
	Amount    *core.Bid `json:"amount"`
}

// Bid auction bid
func (bidReq *BidRequest) Bid(userId string, bid float64) (item *core.Bid, err error) {
	aucRequest := AuctionRequest{}
	auction, err := aucRequest.Get([]byte(bidReq.AuctionId))
	if err != nil {
		return nil, err
	}
	assetKey := []byte(auction.AssetId)
	assetRequest := AssetRequest{}
	asset, err := assetRequest.Get(assetKey)
	if err != nil {
		return nil, err
	}
	if bid < asset.Price {
		err = errors.New("price cannot be lower than the starting price")
		return
	}
	balanceReq := BalanceRequest{}
	balance, err := balanceReq.Get([]byte(userId))
	if err != nil {
		return nil, err
	}
	if bid > balance.Account {
		err = ErrInsufficientFundsBalance
		return
	}
	bidExist, err := bidReq.Exist(userId)
	if !bidExist {
		bidReq.Amount = &core.Bid{
			UserId: userId,
			Bid:    bid,
			Time:   time.Now(),
		}
		return bidReq.Add()
	}
	previousBid, err := bidReq.Get(userId)
	if err != nil {
		return nil, err
	}
	if bid < previousBid.Bid {
		err = errors.New("price cannot be lower than the starting bid")
		return
	}
	bidReq.Amount = &core.Bid{
		UserId: userId,
		Bid:    bid,
		Time:   time.Now(),
	}
	return bidReq.Add()
}

// Add by id
func (bidReq *BidRequest) Add() (item *core.Bid, err error) {
	db, err := cached.OpenDb(AuctionTable)
	if err != nil {
		return
	}
	item = bidReq.Amount
	err = db.BucketAdd([]byte(bidReq.AuctionId), item.Key(), item.Value())
	return
}

func (bidReq *BidRequest) Exist(userId string) (exist bool, err error) {
	db, err := cached.OpenDb(AuctionTable)
	if err != nil {
		return
	}
	auctionId := []byte(bidReq.AuctionId)
	userKey := []byte(userId)
	err, exist = db.BucketExists(auctionId, userKey)
	if err != nil {
		return false, err
	}
	return
}

// Get bid by userId
func (bidReq *BidRequest) Get(userId string) (item *core.Bid, err error) {
	db, err := cached.OpenDb(AuctionTable)
	if err != nil {
		return
	}
	auctionId := []byte(bidReq.AuctionId)
	userKey := []byte(userId)
	err, itemByte := db.BucketGet(auctionId, userKey)
	if err != nil {
		return
	}
	err = json.Unmarshal(itemByte, &item)
	if err != nil {
		return
	}
	return
}

// GetBids get all bids by auction
func (bidReq *BidRequest) GetBids() (items []core.Bid, err error) {
	db, err := cached.OpenDb(AuctionTable)
	if err != nil {
		return
	}
	auctionId := []byte(bidReq.AuctionId)
	values, err := db.BucketGetValues(auctionId)
	if err != nil {
		return
	}
	for _, value := range values {
		var current core.Bid
		err = json.Unmarshal(value.Value, &current)
		if err != nil {
			continue
		}
		items = append(items, current)
	}
	return
}
