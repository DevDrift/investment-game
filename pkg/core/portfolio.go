package core

import "github.com/DevDrift/investment-game/pkg/utils"

// Portfolio user's investment portfolio
type Portfolio struct {
	Id string `json:"id"`
}

// Key - функция для получения ключа.
func (portfolio *Portfolio) Key() []byte {
	return []byte(portfolio.Id)
}

// Value get value bytes
func (portfolio *Portfolio) Value() []byte {
	return utils.ToJsonBytes(portfolio)
}
