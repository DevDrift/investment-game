package models

import (
	"github.com/DevDrift/investment-game/pkg/core"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalanceRequest(t *testing.T) {
	testId := "test"
	req := BalanceRequest{
		Balance: &core.Balance{
			Id:      testId,
			Account: 3.5,
		},
	}
	addBalance, err := req.Add()
	if err != nil {
		t.Error(err)
		return
	}
	getBalance, err := req.Get([]byte(testId))
	if err != nil {
		return
	}
	assert.Equal(t, addBalance.Account, getBalance.Account)
	list, err := req.List()
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, 0, len(list))
	getBalance.Id = uuid.NewString()
	getBalance.Account += 10.0
	req.Balance = getBalance
	updateBalance, err := req.Update([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	updateBalance, err = req.Get([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, updateBalance.Id, testId)
	assert.NotEqual(t, updateBalance.Account, 10.0)
	err = req.Delete([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	deletedBalance, err := req.Get([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, deletedBalance, 0)
}
