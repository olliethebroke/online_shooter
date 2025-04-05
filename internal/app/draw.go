package app

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Draw draws the game screen by one frame.
//
// The give argument represents a screen image. The updated content is adopted as the game screen.
//
// The frequency of Draw calls depends on the user's environment, especially the monitors refresh rate.
// For portability, you should not put your game logic in Draw in general.
func (a *App) Draw(screen *ebiten.Image) {
	// draw menu if it is required
	if a.menu.Active {
		a.menu.Draw(screen)
	}

	// draw game if it is required
	if a.game.Active {
		a.game.Draw(screen)
	}
}
