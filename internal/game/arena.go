package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
	"math/rand"
)

type Arena struct {
	Obstacles []*Obstacle
}

type Obstacle struct {
	Position Point
}

// newArena initializes arena object of the game.
func (g *Game) newArena() {
	g.Arena = &Arena{}
	g.Arena.generateObstacles()
}

// drawArena draws the game's arena objects.
//
// Accepts screen object as an argument.
func (g *Game) drawArena(screen *ebiten.Image) {
	// draw obstacles
	g.Arena.drawObstacles(screen)
}

// drawObstacles draws obstacles of the arena.
//
// Accepts screen object as an argument.
func (a *Arena) drawObstacles(screen *ebiten.Image) {
	// draw every obstacle in a loop
	for _, o := range a.Obstacles {
		vector.DrawFilledRect(
			screen,
			float32(o.Position.X),
			float32(o.Position.Y),
			obstacleSize,
			obstacleSize,
			color.RGBA{R: 139, G: 69, B: 19, A: 255},
			true,
		)
	}
}

// generateObstacles creates arena obstacles
// generating random positions to each obstacle.
func (a *Arena) generateObstacles() {
	// generating numberOfObstacles obstacles
	for len(a.Obstacles) < numberOfObstacles {
		// random coords
		x := rand.Intn(int(0.8*screenWidth)) + int(0.1*screenWidth)
		y := rand.Intn(int(0.8*screenHeight)) + int(0.1*screenHeight)

		// if position is valid
		if a.isPositionValid(x, y) {
			// add an obstacles to the slice
			o := &Obstacle{Position: Point{X: x, Y: y}}
			a.Obstacles = append(a.Obstacles, o)
		}
	}
}

// isPositionValid checks the coords to be valid for
// creating a new obstacle.
//
// Accepts screen coords x and y as a parameters.
//
// Returns true if the position is valid,
// false if it's not.
func (a *Arena) isPositionValid(x, y int) bool {
	// if the obstacle is in the center of the screen
	if math.Abs(float64(x-screenWidth/2)) <= playerSize && math.Abs(float64(y-screenHeight/2)) <= playerSize {
		return false
	}

	// if the obstacle intersects with other obstacles
	for _, obstacle := range a.Obstacles {
		if math.Abs(float64(x-obstacle.Position.X)) < obstacleSize && math.Abs(float64(y-obstacle.Position.Y)) < obstacleSize {
			return false
		}
	}

	// if the obstacle is in the right place
	// return true
	return true
}
