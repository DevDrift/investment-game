package core

import "github.com/DevDrift/investment-game/pkg/utils"

// Player - the structure of the player.
// I'd of the player in the discord.
type Player struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Key get key bytes
func (player *Player) Key() []byte {
	return []byte(player.Id)
}

// Value get value bytes
func (player *Player) Value() []byte {
	return utils.ToJsonBytes(player)
}
