package models

import (
	"encoding/json"
	"github.com/DevDrift/investment-game/pkg/core"
	cached "github.com/DevDrift/investment-game/pkg/database"
)

const PlayerTable = "players"

// PlayerRequest - модель пользователей для взаимодействия с БД.
// Player - объект пользователя.
type PlayerRequest struct {
	Player *core.Player
}

// Add by id
func (req *PlayerRequest) Add() (item *core.Player, err error) {
	db, err := cached.OpenDb(PlayerTable)
	if err != nil {
		return
	}
	item = req.Player
	err = db.BitAdd(item.Key(), item.Value())
	return
}

// Get by id
func (req *PlayerRequest) Get(id []byte) (item *core.Player, err error) {
	db, err := cached.OpenDb(PlayerTable)
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

// List items
func (req *PlayerRequest) List() (items []core.Player, err error) {
	db, err := cached.OpenDb(PlayerTable)
	if err != nil {
		return
	}
	values, err := db.GetValues()
	if err != nil {
		return
	}
	for _, value := range values {
		var current core.Player
		err = json.Unmarshal(value.Value, &current)
		if err != nil {
			continue
		}
		items = append(items, current)
	}
	return
}

// Update by id
func (req *PlayerRequest) Update(id []byte) (item *core.Player, err error) {
	db, err := cached.OpenDb(PlayerTable)
	if err != nil {
		return
	}
	item = req.Player
	err = db.BitAdd(id, item.Value())
	return
}

// Delete by id
func (req *PlayerRequest) Delete(id []byte) (err error) {
	db, err := cached.OpenDb(PlayerTable)
	if err != nil {
		return
	}
	return db.Delete(id)
}

// PutMoney - функция для внесения n-ого значения средств на личный счет игрока.
func (req *PlayerRequest) PutMoney(money float64) (balance *core.Balance, err error) {
	if money < 0 {
		money = 0.0
	}
	player := req.Player
	balanceReq := BalanceRequest{}
	exists, err := balanceReq.Exists(player.Key())
	if err != nil {
		return
	}
	if exists {
		balance, err = balanceReq.Get(player.Key())
		if err != nil {
			return
		}
		balance.Account += money
		balanceReq.Balance = balance
		balance, err = balanceReq.Update(player.Key())
		if err != nil {
			return
		}
		return
	}
	balanceReq.Balance = &core.Balance{
		Id:      player.Id,
		Account: money,
	}
	balance, err = balanceReq.Add()
	return
}

func (req *PlayerRequest) BuyActive(activeId []byte) {

}
