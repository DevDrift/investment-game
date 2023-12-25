package core

// Player - the structure of the player.
// DiscordId - id of the player in the discord.
// AssetStore - Asset list. The Asset's ID.
type Player struct {
	DiscordId  string   `json:"discordId"`
	AssetStore []string `json:"assetStore"`
}

func (player *Player) Key() []byte {
	return []byte(player.DiscordId)
}
