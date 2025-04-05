package arena

import (
	"online_shooter/internal/config"
	"online_shooter/internal/game/geometry"
	"sync"
)

const (
	LowObstaclesAmount    = "low"
	MediumObstaclesAmount = "medium"
	HighObstaclesAmount   = "high"
)

type Arena struct {
	Width           float32
	Height          float32
	SquaresAmount   int
	Spawns          []geometry.Point
	ObstaclesAmount int
	ArenaMutex      sync.RWMutex
	Obstacles       map[int64]*Obstacle
}

// NewArena creates and initializes
// an arena instance with generated obstacles
// and spawn points.
//
// Returns pointer to the created arena.
func NewArena(squaresAmount int, obstaclesAmount string) *Arena {
	// create an arena instance
	// and set basic variables
	arena := &Arena{
		Width:           config.ScreenWidth(),
		Height:          config.ScreenHeight(),
		SquaresAmount:   squaresAmount,
		ObstaclesAmount: parseObstaclesAmount(obstaclesAmount),
	}

	// adapt arena parameters to the squares amount
	arena.adaptToSquaresAmount(squaresAmount)

	// generate obstacles
	arena.generateObstacles()

	// create spawns
	arena.generateSpawns()

	// return pointer to the arena
	return arena
}

// adaptToSquaresAmount adapts the arena for the squares amount.
// Method changes arena's parameters to make it comfortable to play.
//
// Accepts an integer value of the squares amount and a string that shows
// the level of arena filling with obstacles.
func (a *Arena) adaptToSquaresAmount(squaresAmount int) {
	k := float32(squaresAmount) / 4

	// if there are too small amount of players
	// keep the arena sizes default
	if k <= 1 {
		return
	}

	// adapt the arena sizes and the obstacles amount
	// to the squares amount
	a.Width *= k
	a.Height *= k
	a.ObstaclesAmount *= int(k)
}

// parseObstaclesAmount parses the string level of
// the obstacles amount to the integer value.
//
// Accepts a string with level of arena filling with obstacles.
//
// Returns an integer value of the obstacles for the input level.
func parseObstaclesAmount(obstaclesAmount string) int {
	if obstaclesAmount == LowObstaclesAmount {
		return 4
	}
	if obstaclesAmount == MediumObstaclesAmount {
		return 8
	}
	if obstaclesAmount == HighObstaclesAmount {
		return 12
	}
	return 8
}
