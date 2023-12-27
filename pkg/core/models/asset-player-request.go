package models

import "github.com/DevDrift/investment-game/pkg/core"

const AssetPlayersTable = "asset-players"

// AssetPlayerRequest assets for player request model
type AssetPlayerRequest struct {
	Asset  *core.Asset
	Player *core.Player
}
