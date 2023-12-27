package core

// Player - the structure of the player.
// I'd of the player in the discord.
// AssetStore - Asset list. The Asset's ID.
type Player struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	AssetStore []string `json:"assetStore"`
}

// Key get key bytes
func (player *Player) Key() []byte {
	return []byte(player.Id)
}
