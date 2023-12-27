package models

import (
	"encoding/json"
	"github.com/DevDrift/investment-game/pkg/core"
	cached "github.com/DevDrift/investment-game/pkg/database"
	"github.com/DevDrift/investment-game/pkg/utils"
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
	bytes := utils.ToJsonBytes(item)
	err = db.BitAdd(id, bytes)
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
