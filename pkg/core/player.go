package core

// Думал насчет того, чтобы сделать инкапсуляцию локального атрибута UserWallet, но решил, что пока не нужно.

// Player - структура игрока.
// DiscordId - id игрока в дискорде.
// UserWallet - кошелек иргока.
// AssetStore - хранилище имущества.
type Player struct {
	DiscordId  string           `json:"discordId"`
	UserWallet float64          `json:"userWallet"`
	AssetStore map[string]Asset `json:"assetStore"`
}

// PutMoney - функция для пополнения кошелька игрока.
func (player *Player) PutMoney(money float64) {
	player.UserWallet += money
}

// OutPutMoney - функция вывода денег из кошелька игрока.
func (player *Player) OutPutMoney(money float64) float64 {
	if money > player.CheckMoney() {
		//недостаточно средств
		return 0
	} else {
		player.UserWallet -= money
	}
	return money
}

// CheckMoney - функция для получения значения кошелька.
func (player *Player) CheckMoney() float64 {
	return player.UserWallet
}

func (player *Player) PutAssetStore() (err error) {
	return
}

func (player *Player) GetAssetStore() map[string]Asset {
	return player.AssetStore
}
