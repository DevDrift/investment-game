package core

// Balance - сруктура для предоставлении информации об общем балансе игрока.
// DiscordId - Id discord игрока.
// PersonalAccount - личный счет игрока.
// ReserveAccount - резервный счет игрока.
type Balance struct {
	DiscordId       string  `json:"discordId"`
	PersonalAccount float64 `json:"personalAccount"`
	ReserveAccount  float64 `json:"reserveAccount"`
}

// Key - функция для получения ключа.
func (wallet *Balance) Key() []byte {
	return []byte(wallet.DiscordId)
}
