package core

import "github.com/DevDrift/investment-game/pkg/utils"

const (
	StockType          = "stock"
	CryptocurrencyType = "cryptocurrency"
	BuildingType       = "building"
)

// Asset asset data structure
type Asset struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Comment string  `json:"comment"`
	ImgUrl  string  `json:"imgUrl"`
	Type    string  `json:"type"`
	Price   float64 `json:"price"`
	Risk    float64 `json:"risk"`
}

// Key get key bytes
func (asset *Asset) Key() []byte {
	return []byte(asset.Id)
}

// Value get value bytes
func (asset *Asset) Value() []byte {
	return utils.ToJsonBytes(asset)
}
