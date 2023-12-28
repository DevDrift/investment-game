package models

import (
	"encoding/json"
	"errors"
	"github.com/DevDrift/investment-game/pkg/core"
	cached "github.com/DevDrift/investment-game/pkg/database"
)

const BalanceTable = "balances"

var (
	ErrInsufficientFundsBalance = errors.New("ErrInsufficientFundsBalance")
)

// BalanceRequest - модель для взаимодействия с БД.
// Balance - ссылка на общий баланс игрока.
type BalanceRequest struct {
	Balance *core.Balance
}

// Add by id
func (req *BalanceRequest) Add() (item *core.Balance, err error) {
	db, err := cached.OpenDb(BalanceTable)
	if err != nil {
		return
	}
	item = req.Balance
	err = db.BitAdd(item.Key(), item.Value())
	return
}

// Get by id
func (req *BalanceRequest) Get(id []byte) (item *core.Balance, err error) {
	db, err := cached.OpenDb(BalanceTable)
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

// Exists by id
func (req *BalanceRequest) Exists(id []byte) (exist bool, err error) {
	db, err := cached.OpenDb(BalanceTable)
	if err != nil {
		return
	}
	err, exist = db.BitExists(id)
	return
}

// List items
func (req *BalanceRequest) List() (items []core.Balance, err error) {
	db, err := cached.OpenDb(BalanceTable)
	if err != nil {
		return
	}
	values, err := db.GetValues()
	if err != nil {
		return
	}
	for _, value := range values {
		var current core.Balance
		err = json.Unmarshal(value.Value, &current)
		if err != nil {
			continue
		}
		items = append(items, current)
	}
	return
}

// Update by id
func (req *BalanceRequest) Update(id []byte) (item *core.Balance, err error) {
	db, err := cached.OpenDb(BalanceTable)
	if err != nil {
		return
	}
	item = req.Balance
	err = db.BitAdd(id, item.Value())
	return
}

// Delete by id
func (req *BalanceRequest) Delete(id []byte) (err error) {
	db, err := cached.OpenDb(BalanceTable)
	if err != nil {
		return
	}
	return db.Delete(id)
}

// GetPersonalAccount - функция для получения информации о всей сумме личного счета.
func (req *BalanceRequest) GetPersonalAccount(id []byte) (err error, money float64) {
	// TODO Not implemented
	return
}

// GetMoneyFromPersonalAccount - функция для вывода n-ого значения из кошелька игрока.
func (req *BalanceRequest) GetMoneyFromPersonalAccount(id []byte, money float64) (err error, _ float64) {
	// TODO Not implemented
	return
}

// PutMoneyToReserveAccount - функция для внесения n-ого значения средств на резервный счет игрока.
func (req *BalanceRequest) PutMoneyToReserveAccount(money float64) (err error, _ float64) {
	// TODO Not implemented
	return
}

// GetReserveAccount - функция для получения информации о всей сумме резервного счета.
func (req *BalanceRequest) GetReserveAccount() (err error, money float64) {
	// TODO Not implemented
	return
}

// GetMoneyFromReserveAccount - функция для вывода n-ого значения из резервного счета.
func (req *BalanceRequest) GetMoneyFromReserveAccount(money float64) (err error, _ float64) {
	// TODO Not implemented
	return
}
