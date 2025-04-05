package entity

import (
	"online_shooter/internal/game/geometry"
	"sync"
)

type Bullet struct {
	sync.RWMutex
	Position geometry.Point
	Vector   geometry.Vector
	Size     float32
	Speed    float32
	Damage   int32
}

// UpdateBulletPosition updates bullet position
// using its vector and speed.
//
// Accepts a delta time value to correct bullet's speed.
func (b *Bullet) UpdateBulletPosition(deltaTime float32) {
	b.Position.X += b.Vector.X * b.Speed * deltaTime
	b.Position.Y += b.Vector.Y * b.Speed * deltaTime
}
