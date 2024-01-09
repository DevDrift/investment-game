package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuctionRequest_New(t *testing.T) {
	req := AssetRequest{}
	newAsset, err := req.Random()
	if err != nil {
		t.Error(err)
		return
	}
	newAsset.Id = testId
	req.Asset = newAsset
	addAsset, err := req.Add()
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, newAsset.Name, addAsset.Name)
	aucRequest := AuctionRequest{}
	auction, err := aucRequest.New(testId, auctionTestStep)
	if err != nil {
		return
	}
	aucRequest.Auction = auction
	assert.NotEqual(t, auction.Id, testId)
	assert.Equal(t, auction.AssetId, testId)
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
	getAuction, err := aucRequest.Get([]byte(aucFindId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, aucFindId, getAuction.Id)
	getAuction.Bid = 100
	aucRequest.Auction = getAuction
	updateAuction, err := aucRequest.Update([]byte(aucFindId))
	if err != nil {
		t.Error(err)
		return
	}
	getAuction, err = aucRequest.Get([]byte(aucFindId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, updateAuction.Bid, getAuction.Bid)
	err = aucRequest.Delete([]byte(aucFindId))
	if err != nil {
		t.Error(err)
		return
	}
	aucFindId = ""
	list, err = aucRequest.List()
	if err != nil {
		t.Error(err)
		return
	}
	for _, aucItem := range list {
		if aucItem.Id == addAuction.Id {
			aucFindId = aucItem.Id
			break
		}
	}
	assert.Equal(t, aucFindId, "")
}
