package core

const (
	StockType          = "stock"
	CryptocurrencyType = "cryptocurrency"
	BuildingType       = "building"
)

// Asset asset structure
type Asset struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Comment string  `json:"comment"`
	Type    string  `json:"type"`
	Price   float64 `json:"price"`
	Risk    float64 `json:"risk"`
}
