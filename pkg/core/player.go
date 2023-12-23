package core

// Player - структура игрока.
// DiscordId - id игрока в дискорде.
// AssetStore - хранилище имущества.
type Player struct {
	DiscordId  string           `json:"discordId"`
	AssetStore map[string]Asset `json:"assetStore"`
}
