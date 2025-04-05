package drawer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"online_shooter/internal/game/camera"
	"online_shooter/internal/game/entity"
	"online_shooter/internal/game/geometry"
)

// DrawBullets draws the player's bullets on the screen.
//
// Accepts pointers to player, screen, camera objects and a color interface as arguments.
func DrawBullets(player *entity.Square, screen *ebiten.Image, camera *camera.Camera, color color.Color) {
	var inCamPosition geometry.Point
	for _, b := range player.Bullets {
		if b != nil {
			// count bullet's position inside the camera
			inCamPosition = camera.WorldToScreen(b.Position)
			vector.DrawFilledRect(
				screen,
				inCamPosition.X,
				inCamPosition.Y,
				b.Size,
				b.Size,
				color,
				true,
			)
		}
	}
}
