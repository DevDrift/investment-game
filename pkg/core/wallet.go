package core

import "errors"

type Wallet struct {
	DiscordId string  `json:"discordId"`
	Money     float64 `json:"money"`
}

// PutMoney - функция для пополнения кошелька игрока.
func (wallet *Wallet) PutMoney(money float64) {
	wallet.Money += money
}

// OutPutMoney - функция вывода денег из кошелька игрока.
func (wallet *Wallet) OutPutMoney(money float64) (err error, _ float64) {
	if money > wallet.CheckMoney() {
		err = errors.New("Недостаточно средств")
		return err, 0
	} else {
		wallet.Money -= money
	}
	return nil, money
}

// CheckMoney - функция для получения значения кошелька.
func (wallet *Wallet) CheckMoney() float64 {
	return wallet.Money
}
