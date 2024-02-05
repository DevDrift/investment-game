package core

import "github.com/DevDrift/investment-game/pkg/utils"

/*groups*/
const (
	StockType          = "stock"
	CryptocurrencyType = "cryptocurrency"
	BuildingType       = "building"
	FactoryType        = "factory"
)

var (
	//BasePrices prices
	BasePrices = map[string]float64{
		StockType:          10.0,
		CryptocurrencyType: 20.0,
		BuildingType:       30.0,
		FactoryType:        40.0,
	}
	//BaseRisks risks
	BaseRisks = map[string]float64{
		StockType:          1.0,
		CryptocurrencyType: 1.5,
		BuildingType:       2.0,
		FactoryType:        2.5,
	}
	// BaseProfits profits
	BaseProfits = map[string]float64{
		StockType:          0.1,
		CryptocurrencyType: 1.0,
		BuildingType:       1.5,
		FactoryType:        2.0,
	}
	//Types all types
	Types = []string{
		StockType,
		BuildingType,
		CryptocurrencyType,
		FactoryType,
	}
)

// Asset asset data structure
type Asset struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Comment string  `json:"comment"`
	ImgUrl  string  `json:"imgUrl"`
	Type    string  `json:"type"`
	Price   float64 `json:"price"`
	Profit  float64 `json:"profit"`
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
