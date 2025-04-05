package arena

import (
	"github.com/chewxy/math32"
	"math/rand"
	"online_shooter/internal/config"
	"online_shooter/internal/game/entity"
	"online_shooter/internal/game/geometry"
	"sync"
	"time"
)

const secondsToRegenerate = 10

type Obstacle struct {
	sync.RWMutex
	Id         int64
	Position   geometry.Point
	Health     int32
	Size       float32
	Vulnerable bool
}

// generateObstacles creates arena obstacles
// generating random positions to each obstacle.
func (a *Arena) generateObstacles() {
	// init the map
	a.Obstacles = make(map[int64]*Obstacle)

	// setting frames for obstacles' position
	maxWidth := 0.8 * a.Width
	minWidth := 0.1 * a.Width
	maxHeight := 0.8 * a.Height
	minHeight := 0.1 * a.Height

	// generating numberOfObstacles obstacles
	for len(a.Obstacles) < a.ObstaclesAmount {
		// random coords
		x := rand.Float32()*(maxWidth-minWidth) + minWidth
		y := rand.Float32()*(maxHeight-minHeight) + minHeight

		// if position is valid
		if a.isPositionValid(x, y) {
			id := rand.Int63()

			for a.Obstacles[id] != nil {
				id = rand.Int63()
			}

			// create and init a new obstacle instance
			o := &Obstacle{
				Id:         id,
				Position:   geometry.Point{X: x, Y: y},
				Health:     config.ObstacleHealth(),
				Size:       config.ObstacleSize(),
				Vulnerable: true,
			}

			// add an obstacle to the map
			a.Obstacles[id] = o
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
func (a *Arena) isPositionValid(x, y float32) bool {
	playerSize := config.SquareSize()
	obstacleSize := config.ObstacleSize()

	// if the obstacle is in the center of the screen
	if math32.Abs(x-a.Width/2) <= playerSize && math32.Abs(y-a.Height/2) <= playerSize {
		return false
	}

	// if the obstacle intersects with other obstacles
	for _, obstacle := range a.Obstacles {
		if math32.Abs(x-obstacle.Position.X) < obstacleSize*2 && math32.Abs(y-obstacle.Position.Y) < obstacleSize*2 {
			return false
		}
	}

	// if the obstacle is in the right place
	// return true
	return true
}

// GetDamage reduces the health
// and the size of the obstacle that was shot.
//
// Accepts a pointer to the bullet that damaged the obstacle
// and a pointer to the slice of the obstacles.
func (o *Obstacle) GetDamage(b *entity.Bullet) {
	// reduce the obstacle's health
	o.Health -= b.Damage

	// if the obstacle's health
	// is less than 0
	if o.Health <= 0 {
		o.Vulnerable = false

		// start obstacle's regeneration
		go func() {
			o.regenerate()
		}()

		return
	}

	// reduce the obstacle's size
	o.Size -= o.Size * float32(b.Damage) / float32(o.Health) / 2
}

// regenerate starts a timer which end
// restores the obstacle's health and size
// and makes the obstacle visible.
func (o *Obstacle) regenerate() {
	// set a timer
	timer := time.NewTimer(secondsToRegenerate * time.Second)
	defer timer.Stop()

	// wait till the timer ends
	<-timer.C

	o.Lock()
	defer o.Unlock()

	// restore stats
	o.Health = config.ObstacleHealth()
	o.Size = config.ObstacleSize()
	o.Vulnerable = true
}
