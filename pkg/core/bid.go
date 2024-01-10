package core

import (
	"github.com/DevDrift/investment-game/pkg/utils"
	"time"
)

// Bid struct
type Bid struct {
	UserId string    `json:"userId"`
	Bid    float64   `json:"bid"`
	Time   time.Time `json:"time"`
}

// Key get key bytes
func (bid *Bid) Key() []byte {
	return []byte(bid.UserId)
}

// Value get value bytes
func (bid *Bid) Value() []byte {
	return utils.ToJsonBytes(bid)
}
