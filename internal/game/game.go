package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1400
	screenHeight = 800
)

type Game struct {
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
func (g *Game) Draw(screen *ebiten.Image) {}

// Run creates and initializes the game config,
// then starts the game.
//
// Returns the error if it exists, if not - returns nil.
func Run() error {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Shooter")
	g := &Game{}
	if err := ebiten.RunGame(g); err != nil {
		return err
	}
	return nil
}
