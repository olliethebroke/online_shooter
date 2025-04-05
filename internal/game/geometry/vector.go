package geometry

import (
	"github.com/chewxy/math32"
)

type Vector struct {
	X float32
	Y float32
}

// Length returns length of the vector.
func (v *Vector) Length() float32 {
	return math32.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize changes the vector to be its normalized version.
func (v *Vector) Normalize() {
	l := v.Length()
	if l > 0 {
		v.X /= l
		v.Y /= l
	}
}

// VectorToPoint converts the Vector typed var
// to a Point typed var.
//
// Returns a pointer to the created Point object.
func (v *Vector) VectorToPoint() *Point {
	return &Point{
		X: v.X,
		Y: v.Y,
	}
}
