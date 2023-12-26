package models

import (
	"encoding/json"
	"fmt"
	"github.com/DevDrift/investment-game/pkg/core"
	cached "github.com/DevDrift/investment-game/pkg/database"
	"github.com/DevDrift/investment-game/pkg/integration"
	"github.com/DevDrift/investment-game/pkg/utils"
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

// Add asset by id
func (req *AssetRequest) Add() (asset *core.Asset, err error) {
	db, err := cached.OpenDb(AssetsTable)
	if err != nil {
		return
	}
	asset = req.Asset
	bytes := utils.ToJsonBytes(asset)
	err = db.BitAdd(asset.Key(), bytes)
	return
}

// Get asset by id
func (req *AssetRequest) Get(id []byte) (asset *core.Asset, err error) {
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
	err = json.Unmarshal(bytes, &asset)
	if err != nil {
		return
	}
	return
}

// List assets
func (req *AssetRequest) List() (assets []core.Asset, err error) {
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
		assets = append(assets, current)
	}
	return
}

// Update asset by id
func (req *AssetRequest) Update(id []byte) (asset *core.Asset, err error) {
	db, err := cached.OpenDb(AssetsTable)
	if err != nil {
		return
	}
	asset = req.Asset
	bytes := utils.ToJsonBytes(asset)
	err = db.BitAdd(id, bytes)
	return
}

// Delete asset by id
func (req *AssetRequest) Delete(id []byte) (err error) {
	db, err := cached.OpenDb(AssetsTable)
	if err != nil {
		return
	}
	return db.Delete(id)
}
