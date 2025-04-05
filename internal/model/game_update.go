package model

import (
	"online_shooter/internal/game/arena"
	"online_shooter/internal/game/entity"
)

type GameUpdateMessage struct {
	Obstacles map[int64]*arena.Obstacle `json:"obstacles"`
	Squares   map[int64]*entity.Square  `json:"squares"`
}
