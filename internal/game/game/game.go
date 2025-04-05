package game

import (
	"math/rand"
	"online_shooter/internal/game/arena"
	"online_shooter/internal/game/camera"
	"online_shooter/internal/game/entity"
	"online_shooter/internal/menu"
	"sync"
)

type Game struct {
	ricochet  bool
	Active    bool
	GameMutex sync.RWMutex
	Arena     *arena.Arena
	Camera    *camera.Camera
	Squares   map[int64]*entity.Square
	Player    *entity.Square
}

// InitServerGame inits the game with settings parameters.
// Method inits the game arena and generates the required amount of Squares.
//
// Accepts a pointer to the server settings instance.
func (g *Game) InitServerGame(settings *menu.ServerSettings) {
	// init the arena
	g.Arena = arena.NewArena(settings.PlayerCount, settings.ObstacleLevel)

	// generate squares
	g.generateSquares(settings.PlayerCount)

	// set the flag that game is active
	g.Active = true
}

// InitClientGame inits the client game
// setting the game to active state and
// initializing camera with correct position.
func (g *Game) InitClientGame() {
	// init camera
	g.Camera = camera.NewCamera(g.Arena.Width, g.Arena.Height)

	// place camera to the player
	g.Camera.Position = g.Player.Position

	// set the flag that game is active
	g.Active = true
}

// generateSquares fills the Squares slice with
// bot Squares.
//
// Accepts an integer value as an amount of the Squares.
func (g *Game) generateSquares(squaresAmount int) {
	// init the map
	g.Squares = make(map[int64]*entity.Square)

	// generate bots for the rest part of the Squares
	for i := 0; i < squaresAmount; i++ {
		bot := entity.NewBot(g.GenerateUniqueId())
		bot.Position = g.Arena.Spawns[i]
		bot.Spawn = g.Arena.Spawns[i]
		g.Squares[bot.Id] = bot
	}
}

// GenerateUniqueId generates a unique
// identifier for the square.
//
// Returns the generated identifier.
func (g *Game) GenerateUniqueId() int64 {
	var id int64
	for {
		id = rand.Int63()
		exists := false

		// go through the squares
		// to compare ids
		for _, s := range g.Squares {
			if s.Id == id {
				exists = true
				break
			}
		}
		// if such id doesn't exist
		if !exists {
			// leave the infinite loop
			break
		}
	}
	return id
}
