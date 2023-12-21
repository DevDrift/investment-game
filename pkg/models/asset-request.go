package models

import (
	"github.com/DevDrift/investment-game/pkg/core"
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
	}
	min := 0
	max := len(types) - 1
	number := rand.Intn(max-min) + min
	newAsset = &core.Asset{
		Id:      uuid.NewString(),
		Name:    types[number],
		Comment: "",
		Type:    "",
		Price:   0,
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
