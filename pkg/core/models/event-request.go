package models

import (
	"github.com/DevDrift/investment-game/pkg/core"
	"golang.org/x/sync/errgroup"
)

type EventRequest struct {
}

// Bidding by time
func (req *EventRequest) Bidding() (err error) {
	allPlayerReq := PlayerRequest{}
	list, err := allPlayerReq.List()
	if err != nil {
		return
	}
	g := new(errgroup.Group)
	for _, player := range list {
		pl := player
		g.Go(func() (err error) {
			playerKey := pl.Key()
			portfolioRequest := PortfolioRequest{Portfolio: &core.Portfolio{Id: pl.Id}}
			assets, err := portfolioRequest.GetAssets()
			if err != nil {
				return
			}
			var result float64
			for _, asset := range assets {
				// математика
				result += asset.Profit
			}
			balanceReq := BalanceRequest{}
			balance, err := balanceReq.Get(playerKey)
			if err != nil {
				return
			}
			balance.Account += result
			balanceReq.Balance = balance
			_, err = balanceReq.Update(playerKey)
			return
		})
	}
	return g.Wait()
}
