package models

import (
	"github.com/DevDrift/investment-game/pkg/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBidRequest_Bid(t *testing.T) {
	/*balance request*/
	balanceRequest := BalanceRequest{
		Balance: &core.Balance{
			Id:      testId,
			Account: 100,
		},
	}
	addBalance, err := balanceRequest.Add()
	if err != nil {
		t.Error(err)
		return
	}
	getBalance, err := balanceRequest.Get([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, addBalance.Account, getBalance.Account)
	/*asset request*/
	assetRequest := AssetRequest{}
	random, err := assetRequest.Random()
	if err != nil {
		t.Error(err)
		return
	}
	random.Id = testId
	assetRequest.Asset = random
	addAsset, err := assetRequest.Add()
	if err != nil {
		t.Error(err)
		return
	}
	getAsset, err := assetRequest.Get([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, addAsset.Name, getAsset.Name)
	/*auction*/
	aucRequest := AuctionRequest{}
	auction, err := aucRequest.New(getAsset.Id, auctionTestStep)
	if err != nil {
		return
	}
	aucRequest.Auction = auction
	assert.NotEqual(t, auction.Id, testId)
	assert.Equal(t, auction.AssetId, getAsset.Id)
	addAuction, err := aucRequest.Add()
	if err != nil {
		t.Error(err)
		return
	}
	list, err := aucRequest.List()
	if err != nil {
		t.Error(err)
		return
	}
	var aucFindId string
	for _, aucItem := range list {
		if aucItem.Id == addAuction.Id {
			aucFindId = aucItem.Id
			break
		}
	}
	assert.Equal(t, aucFindId, addAuction.Id)
	bidRequest := BidRequest{
		AuctionId: aucFindId,
	}
	/*bids*/
	newBid, err := bidRequest.Bid(testId, 50.0)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, newBid.Bid, 50.0)
	bids, err := bidRequest.GetBids()
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, len(bids), 0)
	_, err = bidRequest.Bid(testId, 40.0)
	assert.NotEqual(t, err, nil)
	newBid, err = bidRequest.Bid(testId, 70.0)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, newBid.Bid, 70.0)
}
