package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPortfolioRequest_BuyAsset(t *testing.T) {
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
	reqBalance := BalanceRequest{
		Balance: testBalance1,
	}
	addBalance, err := reqBalance.Add()
	if err != nil {
		t.Error(err)
		return
	}
	getBalance, err := reqBalance.Get([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, addBalance.Account, getBalance.Account)
	portfolioReq := PortfolioRequest{
		Portfolio: testPortfolio,
	}
	portfolio, err := portfolioReq.BuyAsset(*addAsset)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, portfolio.Id, testPortfolio.Id)
	newBalance, err := reqBalance.Get([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, newBalance.Account, getBalance.Account-newAsset.Price)
	assets, err := portfolioReq.GetAssets()
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, len(assets), 0)
}
