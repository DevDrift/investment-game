package core

// Balance - сруктура для предоставлении информации об общем балансе игрока.
// Id - Id discord игрока.
// Account - личный счет игрока.
type Balance struct {
	Id      string  `json:"id"`
	Account float64 `json:"account"`
}

// Key - функция для получения ключа.
func (wallet *Balance) Key() []byte {
	return []byte(wallet.Id)
}
