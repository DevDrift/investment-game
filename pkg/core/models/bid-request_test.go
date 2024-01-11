package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBidRequest_Bid(t *testing.T) {
	clearAll(t)
	/*balance request*/
	user1 := "user1"
	user2 := "user2"
	user3 := "user3"
	_ = createBalance(t, user1, 300)
	_ = createBalance(t, user2, 300)
	_ = createBalance(t, user3, 300)
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
	user1NewBid, err := bidRequest.Bid(user1, 50.0)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, user1NewBid.Bid, 50.0)
	bids, err := bidRequest.GetBids()
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, len(bids), 0)
	_, err = bidRequest.Bid(user2, 40.0)
	assert.Equal(t, err, nil)
	user1NewBid, err = bidRequest.Bid(user1, 70.0)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, user1NewBid.Bid, 70.0)
	_, err = bidRequest.Bid(user3, 70.0)
	assert.NotEqual(t, err, nil)
	bids, err = bidRequest.GetBids()
	if err != nil {
		t.Error(err)
		return
	}
	for _, bid := range bids {
		t.Logf("U: %v B:%v", bid.UserId, bid.Bid)
	}
}
