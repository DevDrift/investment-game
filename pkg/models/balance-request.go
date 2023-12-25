package models

import "github.com/DevDrift/investment-game/pkg/core"

// BalanceRequest - модель для взаимодействия с БД.
// Balance - ссылка на общий баланс игрока.
type BalanceRequest struct {
	Balance *core.Balance
}

// PutMoneyToPersonalAccount - функция для внесения n-ого значения средств на личный счет игрока.
func (req *BalanceRequest) PutMoneyToPersonalAccount(money float64) (err error) {
	return
}

// GetPersonalAccount - функция для получения информации о всей сумме личного счета.
func (req *BalanceRequest) GetPersonalAccount() (err error, money float64) {
	return
}

// GetMoneyFromPersonalAccount - функция для вывода n-ого значения из кошелька игрока.
func (req *BalanceRequest) GetMoneyFromPersonalAccount(money float64) (err error, _ float64) {
	return
}

// PutMoneyToReserveAccount - функция для внесения n-ого значения средств на резервный счет игрока.
func (req *BalanceRequest) PutMoneyToReserveAccount(money float64) (err error, _ float64) {
	return
}

// GetReserveAccount - функция для получения информации о всей сумме резервного счета.
func (req *BalanceRequest) GetReserveAccount() (err error, money float64) {
	return
}

// GetMoneyFromReserveAccount - функция для вывода n-ого значения из резервного счета.
func (req *BalanceRequest) GetMoneyFromReserveAccount(money float64) (err error, _ float64) {
	return
}
