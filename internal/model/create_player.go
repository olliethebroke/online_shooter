package model

import (
	"online_shooter/internal/game/arena"
	"online_shooter/internal/game/entity"
)

type CreatePlayerResponse struct {
	Arena  *arena.Arena   `json:"arena"`
	Player *entity.Square `json:"player"`
}
