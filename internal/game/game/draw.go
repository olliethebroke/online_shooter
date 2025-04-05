package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	drawer2 "online_shooter/internal/game/drawer"
)

// Draw draws the game on the screen.
//
// Accepts a pointer to the image as an argument.
func (g *Game) Draw(screen *ebiten.Image) {
	g.GameMutex.RLock()
	defer g.GameMutex.RUnlock()

	// draw squares and bullets
	for _, s := range g.Squares {
		drawer2.DrawSquare(s, screen, g.Camera)
		drawer2.DrawBullets(s, screen, g.Camera, s.Color)
	}

	// draw obstacles
	for _, o := range g.Arena.Obstacles {
		drawer2.DrawObstacle(o, screen, g.Camera)
	}

	// draw player's stats
	drawer2.DrawSquareStats(g.Player, screen)
}
