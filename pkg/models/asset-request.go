package models

import (
	"encoding/json"
	"fmt"
	"github.com/DevDrift/investment-game/pkg/core"
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
	types := []string{
		core.StockType,
		core.BuildingType,
		core.CryptocurrencyType,
		core.FactoryType,
	}
	min := 0
	max := len(types) - 1
	number := rand.Intn(max-min) + min
	getType := types[number]
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
		Risk:    0,
	}
	return
}

// Add asset by id
func (req *AssetRequest) Add() (asset *core.Asset, err error) {
	return
}

// Get asset by id
func (req *AssetRequest) Get(id []byte) (asset *core.Asset, err error) {
	return
}

// List assets
func (req *AssetRequest) List() (assets []core.Asset, err error) {
	return
}

// Update asset by id
func (req *AssetRequest) Update(id []byte) (asset *core.Asset, err error) {
	return
}

// Delete asset by id
func (req *AssetRequest) Delete(id []byte) (err error) {
	return
}
