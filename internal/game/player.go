package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Point struct {
	X, Y int
}

type Player struct {
	Position Point
	Color    color.Color
}

// newPlayer creates and initializes
// new player instance of the game.
func (g *Game) newPlayer() {
	p := &Player{
		Color: randomBrightColor(),
		Position: Point{
			X: screenWidth / 2,
			Y: screenHeight / 2,
		},
	}
	g.Players = append(g.Players, p)
}

// drawPlayers draws players of the game.
//
// Accepts screen object as an argument.
func (g *Game) drawPlayers(screen *ebiten.Image) {
	for _, p := range g.Players {
		vector.DrawFilledRect(
			screen,
			float32(p.Position.X),
			float32(p.Position.Y),
			playerSize,
			playerSize,
			p.Color,
			true,
		)
	}
}
