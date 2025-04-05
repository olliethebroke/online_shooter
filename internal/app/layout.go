package app

import "online_shooter/internal/config"

// Layout accepts a native outside size in device-independent pixels
// and returns the game's logical screen size in pixels.
func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(config.ScreenWidth()), int(config.ScreenHeight())
}
