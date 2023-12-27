package models

import (
	"encoding/json"
	"fmt"
	"github.com/DevDrift/investment-game/pkg/core"
	cached "github.com/DevDrift/investment-game/pkg/database"
	"github.com/DevDrift/investment-game/pkg/integration"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

const AssetsTable = "assets"

// AssetRequest asset request model
type AssetRequest struct {
	Asset *core.Asset
}

// Random create random asset
func (req *AssetRequest) Random() (newAsset *core.Asset, err error) {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(core.Types) - 1
	number := rand.Intn(max-min) + min
	getType := core.Types[number]
	/*names*/
	typeDataNameRequest := integration.JsonRequest{
		Url: fmt.Sprintf("https://raw.githubusercontent.com/DevDrift/investment-game/main/data/names/%s.json", getType),
	}
	dataNames, err := typeDataNameRequest.GetRawData()
	if err != nil {
		return nil, err
	}
	var names []string
	err = json.Unmarshal(dataNames, &names)
	if err != nil {
		return
	}
	max = len(names) - 1
	randNumberName := rand.Intn(max-min) + min
	/*prompts*/
	typeDataPromptsRequest := integration.JsonRequest{
		Url: fmt.Sprintf("https://raw.githubusercontent.com/DevDrift/investment-game/main/data/prompts/%s.json", getType),
	}
	dataPrompts, err := typeDataPromptsRequest.GetRawData()
	if err != nil {
		return
	}
	var prompts []string
	err = json.Unmarshal(dataPrompts, &prompts)
	if err != nil {
		return
	}
	max = len(prompts) - 1
	randNumberPrompt := rand.Intn(max-min) + min
	max = 10
	randNumberImage := rand.Intn(max-min) + min
	newAsset = &core.Asset{
		Id:      uuid.NewString(),
		Name:    names[randNumberName],
		Comment: prompts[randNumberPrompt],
		Type:    getType,
		ImgUrl:  fmt.Sprintf("https://raw.githubusercontent.com/DevDrift/investment-game/images/%s/%d.png", getType, randNumberImage),
		Price:   core.BasePrices[getType],
		Risk:    core.BaseRisks[getType],
	}
	return
}

// Add by id
func (req *AssetRequest) Add() (item *core.Asset, err error) {
	db, err := cached.OpenDb(AssetsTable)
	if err != nil {
		return
	}
	item = req.Asset
	err = db.BitAdd(item.Key(), item.Value())
	return
}

// Get by id
func (req *AssetRequest) Get(id []byte) (item *core.Asset, err error) {
	db, err := cached.OpenDb(AssetsTable)
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
func (req *AssetRequest) List() (items []core.Asset, err error) {
	db, err := cached.OpenDb(AssetsTable)
	if err != nil {
		return
	}
	values, err := db.GetValues()
	if err != nil {
		return
	}
	for _, value := range values {
		var current core.Asset
		err = json.Unmarshal(value.Value, &current)
		if err != nil {
			continue
		}
		items = append(items, current)
	}
	return
}

// Update by id
func (req *AssetRequest) Update(id []byte) (item *core.Asset, err error) {
	db, err := cached.OpenDb(AssetsTable)
	if err != nil {
		return
	}
	item = req.Asset
	err = db.BitAdd(id, item.Value())
	return
}

// Delete by id
func (req *AssetRequest) Delete(id []byte) (err error) {
	db, err := cached.OpenDb(AssetsTable)
	if err != nil {
		return
	}
	return db.Delete(id)
}
