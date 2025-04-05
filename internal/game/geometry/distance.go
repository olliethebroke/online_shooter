package geometry

import (
	"github.com/chewxy/math32"
)

// GetDistanceBetweenTwoPoints counts a distance in pixels
// between two points.
//
// Accepts two Point objects as arguments.
//
// Returns the distance.
func GetDistanceBetweenTwoPoints(p1, p2 Point) float32 {
	// count differences between positions
	diffX := p1.X - p2.X
	diffY := p1.Y - p2.Y

	// return the distance
	return math32.Hypot(diffX, diffY)
}
