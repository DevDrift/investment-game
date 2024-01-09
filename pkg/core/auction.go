package core

import (
	"github.com/DevDrift/investment-game/pkg/utils"
	"time"
)

// Auction auction data structure
type Auction struct {
	Id        string    `json:"id"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	AssetId   string    `json:"assetId"`
	Bid       float64   `json:"bid"`
	Step      float64   `json:"step"`
	WinnerId  string    `json:"winnerId"`
}

// Key get key bytes
func (auc *Auction) Key() []byte {
	return []byte(auc.Id)
}

// Value get value bytes
func (auc *Auction) Value() []byte {
	return utils.ToJsonBytes(auc)
}
