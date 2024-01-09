package models

import (
	"encoding/json"
	"github.com/DevDrift/investment-game/pkg/core"
	cached "github.com/DevDrift/investment-game/pkg/database"
)

const PortfolioTable = "portfolios"

type PortfolioRequest struct {
	Portfolio *core.Portfolio
}

// BuyAsset buy asset
func (req *PortfolioRequest) BuyAsset(asset core.Asset) (item *core.Portfolio, err error) {
	db, err := cached.OpenDb(PortfolioTable)
	if err != nil {
		return
	}
	item = req.Portfolio
	userId := item.Key()
	assetPrice := asset.Price
	balanceReq := BalanceRequest{}
	balance, err := balanceReq.Get(userId)
	if err != nil {
		return
	}
	if balance.Account < assetPrice {
		return nil, ErrInsufficientFundsBalance
	}
	err = db.BucketAdd(userId, asset.Key(), asset.Value())
	if err != nil {
		return
	}
	balance.Account -= assetPrice
	balanceReq.Balance = balance
	_, err = balanceReq.Update(userId)
	if err != nil {
		return
	}
	return
}

// GetAssets get list assets from portfolio
func (req *PortfolioRequest) GetAssets() (items []core.Asset, err error) {
	db, err := cached.OpenDb(PortfolioTable)
	if err != nil {
		return
	}
	item := req.Portfolio
	userId := item.Key()
	values, err := db.BucketGetValues(userId)
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

// sell to user
// create auction store
// add asset to auction store
