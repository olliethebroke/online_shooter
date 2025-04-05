package game

import (
	"github.com/chewxy/math32"
	"online_shooter/internal/game/arena"
	"online_shooter/internal/game/entity"
	"online_shooter/internal/game/geometry"
)

// CheckSquareCollision counts the position
// of the square due to the collisions with obstacles.
//
// Accepts Square pointer as an argument.
func (g *Game) CheckSquareCollision(p *entity.Square) {
	p.Lock()
	defer p.Unlock()

	// check if it is a collision with the obstacles
	collision, obstacle := g.checkCollisionWithObstacles(p.Position, p.Size)
	if collision {
		// change square's position
		// to not go through the obstacle
		obstacle.RLock()
		squareCollisionWithObstacle(p, obstacle)
		obstacle.RUnlock()
	}

	// check if it is a collision with other squares
	var square *entity.Square
	collision, square = g.checkCollisionWithSquares(p.Position, p.Size, p)
	if collision {
		// change square's position
		// to not go through the player
		square.RLock()
		squareCollisionWithSquare(p, square)
		square.RUnlock()
	}

	// check if it is a collision with the boards
	g.checkCollisionWithBorders(&p.Position, p.Size)

}

// CheckBulletsCollision checks every bullet collision
// with players, obstacles and borders.
// If there is a collision bullet is removed from the screen.
//
// Accepts a pointer to the player that shot the bullets.
func (g *Game) CheckBulletsCollision(p *entity.Square) {
	p.RLock()
	defer p.RUnlock()

	for i, b := range p.Bullets {
		if b != nil {
			b.Lock()

			// check if it is a collision with obstacles
			collision, obstacle := g.checkCollisionWithObstacles(b.Position, b.Size)
			if collision {
				obstacle.Lock()

				// process the consequences of the obstacle and bullet collision
				obstacle.GetDamage(b)

				// check if ricochet is enabled
				if g.ricochet {
					// if it is - change bullet's vector
					ricochetBullet(b, obstacle)
				} else {
					// otherwise remove the bullet from the Arena
					p.RemoveBullet(i)
					obstacle.Unlock()
					b.Unlock()
					continue
				}
				obstacle.Unlock()
			}

			// check if it is a collision with players
			var damagedPlayer *entity.Square
			collision, damagedPlayer = g.checkCollisionWithSquares(b.Position, b.Size, p)
			if collision {
				// process the consequences of the square and bullet collision
				damagedPlayer.Lock()
				damagedPlayer.GetDamage(b, p)
				damagedPlayer.Unlock()

				// remove the bullet from the arena
				p.RemoveBullet(i)
				b.Unlock()
				continue
			}

			// check if it is a collision with borders
			collision = g.checkCollisionWithBorders(&b.Position, b.Size)
			if collision {
				// remove the bullet from the arena
				p.RemoveBullet(i)
				b.Unlock()
				continue
			}

			b.Unlock()
		}
	}
}

// checkCollisionWithObstacles checks if it is a collision between
// an object and the obstacles.
//
// Accepts the object's position and a size of the object.
//
// Returns true and a pointer to the obstacle which has a collision with the object.
// if there is no collision - method returns false and nil.
func (g *Game) checkCollisionWithObstacles(objectPosition geometry.Point, objectSize float32) (bool, *arena.Obstacle) {
	// go through every obstacle
	g.Arena.ArenaMutex.Lock()
	defer g.Arena.ArenaMutex.Unlock()
	for _, o := range g.Arena.Obstacles {
		o.Lock()

		// skip if the obstacle is invulnerable
		if !o.Vulnerable {
			o.Unlock()
			continue
		}

		// check the collision between the object
		// and the obstacle
		if isCollision(objectPosition, o.Position, objectSize, o.Size) {
			// if there is a collision
			o.Unlock()
			return true, o
		}

		o.Unlock()
	}

	// if there is no collision
	return false, nil
}

// checkCollisionWithSquares checks if it is a collision between
// an object and the squares.
//
// Accepts the object's position, a size of the object
// and a pointer to the square that checks a collision.
//
// Returns true and a pointer to the square which has a collision with the object.
// if there is no collision - method returns false and nil.
func (g *Game) checkCollisionWithSquares(objectPosition geometry.Point, objectSize float32, shooter *entity.Square) (bool, *entity.Square) {
	// go through every square
	for _, s := range g.Squares {
		// square to check collision must be different from the square that asks for check
		if shooter == s {
			continue
		}

		s.RLock()

		// if the square to check collision is invulnerable
		// skip the check
		if !s.Vulnerable {
			s.RUnlock()
			continue
		}

		// check the collision between the object
		// and the square
		if isCollision(objectPosition, s.Position, objectSize, s.Size) {
			// if there is a collision
			s.RUnlock()
			return true, s
		}

		s.RUnlock()
	}

	// if there is no collision
	return false, nil
}

// checkCollisionWithBorders checks if it is a collision between
// an object and the borders.
//
// Accepts a pointer to the object's position and a size of the object.
//
// Returns true if there is a collision otherwise returns false.
func (g *Game) checkCollisionWithBorders(objectPosition *geometry.Point, objectSize float32) bool {
	// create a flag that indicates if the object
	// has gone over the borders
	isOverBorder := false
	// check if object goes over the right border
	if objectPosition.X+objectSize > g.Arena.Width {
		objectPosition.X = g.Arena.Width - objectSize
		isOverBorder = true
	}
	// check if object goes over the bottom border
	if objectPosition.Y+objectSize > g.Arena.Height {
		objectPosition.Y = g.Arena.Height - objectSize
		isOverBorder = true
	}
	// check if object goes over the left border
	if objectPosition.X < 0 {
		objectPosition.X = 0
		isOverBorder = true
	}
	// check if object goes over the top border
	if objectPosition.Y < 0 {
		objectPosition.Y = 0
		isOverBorder = true
	}

	return isOverBorder
}

// isCollision check if it is a collision between two objects.
//
// Accepts two pointers to the objects' positions
// and the objects' sizes.
func isCollision(p1, p2 geometry.Point, s1, s2 float32) bool {
	if p1.X < p2.X {
		if p1.X+s1 < p2.X {
			return false
		}
	} else {
		if p2.X+s2 < p1.X {
			return false
		}
	}

	if p1.Y < p2.Y {
		if p1.Y+s1 < p2.Y {
			return false
		}
	} else {
		if p2.Y+s2 < p1.Y {
			return false
		}
	}
	return true
}

// squareCollisionWithObstacle corrects square's position near the obstacle.
//
// Accepts a pointer to the square and a pointer to the obstacle.
func squareCollisionWithObstacle(s *entity.Square, o *arena.Obstacle) {
	// count centers
	playerCenterX := s.Position.X + s.Size/2
	playerCenterY := s.Position.Y + s.Size/2
	obstacleCenterX := o.Position.X + o.Size/2
	obstacleCenterY := o.Position.Y + o.Size/2

	// count differences between centers
	diffX := playerCenterX - obstacleCenterX
	diffY := playerCenterY - obstacleCenterY

	// get the collision side
	if math32.Abs(diffX) > math32.Abs(diffY) {
		// horizontal collision
		if diffX >= 0 {
			// right side
			s.Position.X = o.Position.X + o.Size
		} else {
			// left side
			s.Position.X = o.Position.X - s.Size
		}
	} else {
		// vertical collision
		if diffY >= 0 {
			// bottom side
			s.Position.Y = o.Position.Y + o.Size
		} else {
			// top side
			s.Position.Y = o.Position.Y - s.Size
		}
	}
}

// squareCollisionWithSquare corrects square's position near the square.
//
// Accepts two pointers to the square objects.
func squareCollisionWithSquare(s1, s2 *entity.Square) {
	// count centers
	player1CenterX := s1.Position.X + s1.Size/2
	player1CenterY := s1.Position.Y + s1.Size/2
	player2CenterX := s2.Position.X + s2.Size/2
	player2CenterY := s2.Position.Y + s2.Size/2

	// count differences between centers
	diffX := player1CenterX - player2CenterX
	diffY := player1CenterY - player2CenterY

	// get the collision side
	if math32.Abs(diffX) > math32.Abs(diffY) {
		// horizontal collision
		if diffX >= 0 {
			// right side
			s1.Position.X = s2.Position.X + s2.Size
		} else {
			// left side
			s1.Position.X = s2.Position.X - s1.Size
		}
	} else {
		// vertical collision
		if diffY >= 0 {
			// bottom side
			s1.Position.Y = s2.Position.Y + s2.Size
		} else {
			// top side
			s1.Position.Y = s2.Position.Y - s1.Size
		}
	}
}

// ricochetBullet changes vector of the bullet
// depending on the collision side.
//
// Accepts a pointer to the bullet
// and a pointer to the obstacle.
func ricochetBullet(b *entity.Bullet, o *arena.Obstacle) {
	// count centers
	bulletCenterX := b.Position.X + b.Size/2
	bulletCenterY := b.Position.Y + b.Size/2
	obstacleCenterX := o.Position.X + o.Size/2
	obstacleCenterY := o.Position.Y + o.Size/2

	// count differences between centers
	diffX := bulletCenterX - obstacleCenterX
	diffY := bulletCenterY - obstacleCenterY

	// get the collision side
	if math32.Abs(diffX) > math32.Abs(diffY) {
		// horizontal collision
		if diffX >= 0 {
			// right side
			b.Position.X = o.Position.X + o.Size
		} else {
			// left side
			b.Position.X = o.Position.X - b.Size
		}

		// reverse vector
		b.Vector.X *= -1
	} else {
		// vertical collision
		if diffY >= 0 {
			// bottom side
			b.Position.Y = o.Position.Y + o.Size
		} else {
			// top side
			b.Position.Y = o.Position.Y - b.Size
		}

		// reverse vector
		b.Vector.Y *= -1
	}
}
