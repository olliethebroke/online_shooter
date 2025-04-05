package camera

import (
	"online_shooter/internal/config"
	"online_shooter/internal/game/geometry"
)

type Camera struct {
	Width, Height           float32
	ArenaWidth, ArenaHeight float32
	Position                geometry.Point
}

// NewCamera creates and inits new camera instance.
//
// Accepts arena width and arena height values as arguments.
//
// Returns a pointer to the created camera object.
func NewCamera(arenaWidth, arenaHeight float32) *Camera {
	return &Camera{
		Width:       config.ScreenWidth(),
		Height:      config.ScreenHeight(),
		ArenaWidth:  arenaWidth,
		ArenaHeight: arenaHeight,
	}
}

// Move moves the camera to follow the specified point.
//
// Accepts the point as an argument.
func (c *Camera) Move(point geometry.Point) {
	// count the camera center point
	cameraCenter := geometry.Point{
		X: c.Position.X + c.Width/2,
		Y: c.Position.Y + c.Height/2,
	}

	// get the vector between the center of the camera
	// and the specified point
	vector := geometry.Vector{
		X: point.X - cameraCenter.X,
		Y: point.Y - cameraCenter.Y,
	}

	// get the distance between the camera center point
	// and the specified point
	distance := vector.Length()

	var speed float32
	// if the point is not near the camera center
	if distance > c.Width/2 {
		// make a smooth camera transferring
		speed = 0.05 * distance
	} else {
		// otherwise keep the camera on specified point
		speed = distance
	}

	// normalize the vector
	vector.Normalize()

	// transfer the camera
	c.Position.X += vector.X * speed
	c.Position.Y += vector.Y * speed

	// check the borders
	c.checkBordersCollision()
}

// checkBordersCollision checks the camera's collision
// with the arena borders. If there is a collision
// camera changes its position.
func (c *Camera) checkBordersCollision() {
	if c.Position.X < 0 {
		c.Position.X = 0
	}
	if c.Position.Y < 0 {
		c.Position.Y = 0
	}
	if c.Position.X > c.ArenaWidth-c.Width {
		c.Position.X = c.ArenaWidth - c.Width
	}
	if c.Position.Y > c.ArenaHeight-c.Height {
		c.Position.Y = c.ArenaHeight - c.Height
	}
}

// WorldToScreen converts a point from the game world coordinate
// system to a point from the screen coordinate system.
//
// Accepts the point from the game world coordinate system.
//
// Returns the point from the screen coordinate system.
func (c *Camera) WorldToScreen(point geometry.Point) geometry.Point {
	point.X -= c.Position.X
	point.Y -= c.Position.Y
	return point
}

// ScreenToWorld converts a point from the screen coordinate system
// to a point from the game world coordinate system.
//
// Accepts the point from the screen coordinate system.
//
// Returns the point from the game world coordinate system.
func (c *Camera) ScreenToWorld(point geometry.Point) geometry.Point {
	point.X += c.Position.X
	point.Y += c.Position.Y
	return point
}
