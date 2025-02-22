package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth       = 1400
	screenHeight      = 800
	playerSize        = 25
	obstacleSize      = 40
	numberOfObstacles = 6
)

type Game struct {
	Arena   *Arena
	Players []*Player
}

var game *Game

// init initializes the game.
func init() {
	game = &Game{}
	game.newArena()
	game.newPlayer()
}

// Update updates a game by one tick. The given argument represents a screen image.
//
// Update updates only the game logic and Draw draws the screen.
func (g *Game) Update() error {
	return nil
}

// Layout accepts a native outside size in device-independent pixels
// and returns the game's logical screen size in pixels.
func (g *Game) Layout(outsideWidth, outsideHEight int) (int, int) {
	return screenWidth, screenHeight
}

// Draw draws the game screen by one frame.
//
// The give argument represents a screen image. The updated content is adopted as the game screen.
//
// The frequency of Draw calls depends on the user's environment, especially the monitors refresh rate.
// For portability, you should not put your game logic in Draw in general.
func (g *Game) Draw(screen *ebiten.Image) {
	g.drawPlayers(screen)
	g.drawArena(screen)
}

// Run creates and initializes the game config,
// then starts the game.
//
// Returns the error if it exists, if not - returns nil.
func Run() error {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Shooter")
	if err := ebiten.RunGame(game); err != nil {
		return err
	}
	return nil
}
