package drawer

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"online_shooter/internal/assets"
	"online_shooter/internal/game/camera"
	"online_shooter/internal/game/entity"
)

// DrawSquare draws the square on the screen.
//
// Accepts pointers to the square to draw,
// screen and camera objects as arguments.
func DrawSquare(square *entity.Square, screen *ebiten.Image, camera *camera.Camera) {
	// count player's position inside camera
	inCamPosition := camera.WorldToScreen(square.Position)
	vector.DrawFilledRect(
		screen,
		inCamPosition.X,
		inCamPosition.Y,
		square.Size,
		square.Size,
		square.Color,
		true,
	)
}

// DrawSquareStats draws the statistic about the square on the screen.
// Statistic includes a square's health, an amount of kills and of deaths.
//
// Accepts pointers to the square to draw stats and screen objects as arguments.
func DrawSquareStats(square *entity.Square, screen *ebiten.Image) {
	health := fmt.Sprintf("HEALTH: %d", square.Health)
	bullets := fmt.Sprintf("BULLETS: %d", square.CountBulletsAmount())
	kills := fmt.Sprintf("KILLS: %d", square.Kills)
	deaths := fmt.Sprintf("DEATHS: %d", square.Deaths)

	textColor := color.White

	text.Draw(screen, health, assets.Font(), 10, 30, textColor)
	text.Draw(screen, bullets, assets.Font(), 10, 60, textColor)
	text.Draw(screen, kills, assets.Font(), 10, 90, textColor)
	text.Draw(screen, deaths, assets.Font(), 10, 120, textColor)
}
