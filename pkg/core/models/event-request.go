package models

import "github.com/DevDrift/investment-game/pkg/core"

type EventRequest struct {
}

// Bidding by
func (req *EventRequest) Bidding() (err error) {
	allPlayerReq := PlayerRequest{}
	list, err := allPlayerReq.List()
	if err != nil {
		return
	}
	for _, player := range list {
		playerKey := player.Key()
		portfolioRequest := PortfolioRequest{Portfolio: &core.Portfolio{Id: player.Id}}
		assets, err := portfolioRequest.GetAssets()
		if err != nil {
			return err
		}
		var result float64
		for _, asset := range assets {
			// математика
			result += asset.Profit
		}
		balanceReq := BalanceRequest{}
		balance, err := balanceReq.Get(playerKey)
		if err != nil {
			return err
		}
		balance.Account += result
		balanceReq.Balance = balance
		_, err = balanceReq.Update(playerKey)
		if err != nil {
			return err
		}
	}
	return
}
