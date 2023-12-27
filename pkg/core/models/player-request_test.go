package models

import (
	"github.com/DevDrift/investment-game/pkg/core"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayerRequest(t *testing.T) {
	testId := "test"
	testName := "tName"
	testEmail := "tEmail"
	req := PlayerRequest{
		Player: &core.Player{
			Id:    testId,
			Name:  testName,
			Email: testEmail,
		},
	}
	addPlayer, err := req.Add()
	if err != nil {
		t.Error(err)
		return
	}
	getPlayer, err := req.Get([]byte(testId))
	if err != nil {
		return
	}
	assert.Equal(t, addPlayer.Name, getPlayer.Name)
	list, err := req.List()
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, 0, len(list))
	getPlayer.Id = uuid.NewString()
	getPlayer.Name = "newName"
	req.Player = getPlayer
	updatePlayer, err := req.Update([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	updatePlayer, err = req.Get([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, updatePlayer.Id, testId)
	assert.NotEqual(t, updatePlayer.Name, "")
	err = req.Delete([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	deletedPlayer, err := req.Get([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, deletedPlayer, 0)
}

func TestPlayerRequest_PutMoney(t *testing.T) {
	testId := uuid.NewString()
	testName := "tName"
	testEmail := "tEmail"
	inputMoney := 100.0
	req := PlayerRequest{
		Player: &core.Player{
			Id:    testId,
			Name:  testName,
			Email: testEmail,
		},
	}
	addPlayer, err := req.Add()
	if err != nil {
		t.Error(err)
		return
	}
	getPlayer, err := req.Get([]byte(testId))
	if err != nil {
		return
	}
	assert.Equal(t, addPlayer.Name, getPlayer.Name)
	balance, err := req.PutMoney(inputMoney)
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, balance, 0)
	balanceReq := BalanceRequest{balance}
	getBalance, err := balanceReq.Get([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, inputMoney, getBalance.Account)
	balance, err = req.PutMoney(inputMoney)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, inputMoney*2, balance.Account)
	getBalance, err = balanceReq.Get([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, inputMoney*2, getBalance.Account)
	balance, err = req.PutMoney(-inputMoney)
	if err != nil {
		t.Error(err)
		return
	}
	getBalance, err = balanceReq.Get([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, inputMoney*2, getBalance.Account)
}
