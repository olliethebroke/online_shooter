package entity

import (
	"online_shooter/internal/config"
	"online_shooter/internal/utils"
)

const msToReload = 500

// NewBot creates and initializes
// new bot square instance with default parameters.
//
// Accepts an id as an argument.
//
// Returns pointer to the created bot square.
func NewBot(id int64) *Square {
	// create and init instance
	b := &Square{
		Id:         id,
		Health:     100,
		Speed:      config.SquareSpeed(),
		Size:       config.SquareSize(),
		Color:      utils.RandomBrightColor(),
		CanShoot:   true,
		Vulnerable: true,
		IsBot:      true,
		ShotCh:     make(chan struct{}),
	}

	return b
}
