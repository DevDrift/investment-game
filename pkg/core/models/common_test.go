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
