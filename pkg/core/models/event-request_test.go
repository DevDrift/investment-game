package models

import (
	"github.com/DevDrift/investment-game/pkg/core"
	"github.com/robfig/cron"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEventRequestMinute(t *testing.T) {
	endChannel := make(chan struct{})
	c := cron.New()
	err := c.AddFunc("59 */2 * * * *", func() {
		t.Logf("[%v] event 59 second - run", time.Now())
		endChannel <- struct{}{}
	})
	if err != nil {
		t.Error(err)
		return
	}
	c.Start()
	<-endChannel
}

func TestEventRequest_Bidding(t *testing.T) {
	clearAll(t)
	assets, err := newAssets(t, 10)
	if err != nil {
		t.Error(err)
		return
	}
	user1 := "user1"
	_, err = newPlayer(t, user1)
	if err != nil {
		t.Error(err)
		return
	}
	user2 := "user2"
	_, err = newPlayer(t, user2)
	if err != nil {
		t.Error(err)
		return
	}
	_ = createBalance(t, user1, 3000)
	_ = createBalance(t, user2, 3000)
	userPortfolio1 := PortfolioRequest{Portfolio: &core.Portfolio{Id: user1}}
	userPortfolio2 := PortfolioRequest{Portfolio: &core.Portfolio{Id: user2}}
	for i, asset := range assets {
		if i%2 == 0 {
			userPortfolio2.BuyAsset(asset)
			continue
		}
		userPortfolio1.BuyAsset(asset)
	}
	balanceReq := BalanceRequest{}
	firstBalance1, err := balanceReq.Get([]byte(user1))
	if err != nil {
		t.Error(err)
		return
	}
	firstBalance2, err := balanceReq.Get([]byte(user2))
	if err != nil {
		t.Error(err)
		return
	}
	getAssets1, err := userPortfolio1.GetAssets()
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, len(getAssets1), 5)
	getAssets2, err := userPortfolio2.GetAssets()
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, len(getAssets2), 5)
	eventReq := EventRequest{
		Events: []core.Event{
			{
				Type:    core.CryptocurrencyType,
				Percent: 1.0,
			},
		},
	}
	err = eventReq.Bidding()
	if err != nil {
		t.Error(err)
		return
	}
	getBalance1, err := balanceReq.Get([]byte(user1))
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, firstBalance1.Account, getBalance1.Account)
	getBalance2, err := balanceReq.Get([]byte(user2))
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, firstBalance2.Account, getBalance2.Account)
}
