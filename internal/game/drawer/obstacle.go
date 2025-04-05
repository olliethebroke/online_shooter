package drawer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"online_shooter/internal/game/arena"
	"online_shooter/internal/game/camera"
)

// DrawObstacle draws the obstacle object on the screen.
//
// Accepts pointers to obstacle to draw, screen and camera objects as arguments.
func DrawObstacle(obstacle *arena.Obstacle, screen *ebiten.Image, camera *camera.Camera) {
	// check if it is required to draw the obstacle
	if !obstacle.Vulnerable {
		return
	}

	// count obstacle's position inside camera
	inCamPosition := camera.WorldToScreen(obstacle.Position)
	vector.DrawFilledRect(
		screen,
		inCamPosition.X,
		inCamPosition.Y,
		obstacle.Size,
		obstacle.Size,
		color.RGBA{R: 139, G: 69, B: 19, A: 255},
		true,
	)
}
