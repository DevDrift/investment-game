package models

import (
	"github.com/DevDrift/investment-game/pkg/core"
	"github.com/DevDrift/investment-game/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testId      = "test"
	testName    = "tName"
	testNewName = "newName"
	testEmail   = "tEmail"
	inputMoney  = 100.0
	testPlayer1 = &core.Player{
		Id:    testId,
		Name:  testName,
		Email: testEmail,
	}
	testBalance1 = &core.Balance{
		Id:      testId,
		Account: inputMoney,
	}
	testPortfolio = &core.Portfolio{
		Id: testId,
	}
	auctionTestStep = 0.10
)

func newAssets(t *testing.T, count int) (list []core.Asset, err error) {
	for i := 1; i <= count; i++ {
		req := AssetRequest{}
		newAsset, err := req.Random()
		if err != nil {
			return nil, err
		}
		list = append(list, *newAsset)
	}
	return
}

func newPlayer(t *testing.T, id string) (player *core.Player, err error) {
	req := PlayerRequest{
		Player: &core.Player{
			Id:   id,
			Name: id,
		},
	}
	return req.Add()
}

func createBalance(t *testing.T, userId string, amount float64) (balance *core.Balance) {
	balanceRequest := BalanceRequest{
		Balance: &core.Balance{
			Id:      userId,
			Account: amount,
		},
	}
	balance, err := balanceRequest.Add()
	if err != nil {
		t.Error(err)
		return
	}
	getBalance, err := balanceRequest.Get([]byte(userId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, balance.Account, getBalance.Account)
	return
}

func clearAll(t *testing.T) {
	err := utils.DeleteFilesByExtension("ig-base", ".igb")
	if err != nil {
		t.Error(err)
		return
	}
}
