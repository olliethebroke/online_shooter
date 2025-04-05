package entity

import (
	"context"
	"github.com/gorilla/websocket"
	"image/color"
	"online_shooter/internal/config"
	"online_shooter/internal/game/geometry"
	"online_shooter/internal/utils"
	"sync"
	"time"
)

const (
	bulletsAmount          = 3
	changeColorMs          = 250
	invulnerabilitySeconds = 3
)

type Square struct {
	Id           int64           `json:"id"`
	Conn         *websocket.Conn `json:"-"`
	sync.RWMutex `json:"-"`
	Position     geometry.Point         `json:"position"`
	Spawn        geometry.Point         `json:"-"`
	Health       int32                  `json:"health"`
	Speed        float32                `json:"speed"`
	Size         float32                `json:"size"`
	Bullets      [bulletsAmount]*Bullet `json:"bullets"`
	Kills        uint16                 `json:"kills"`
	Deaths       uint16                 `json:"deaths"`
	Vulnerable   bool                   `json:"-"`
	IsBot        bool                   `json:"is_bot"`
	CanShoot     bool                   `json:"-"`
	Color        color.RGBA             `json:"color"`
	LastUpdate   *time.Time             `json:"-"`
	ShotCh       chan struct{}          `json:"-"`
}

// Move changes the position of the square due to
// the moving vector.
//
// Accepts the moving vector and a delta time correction value.
func (s *Square) Move(v geometry.Vector, deltaTime float32) {
	// count new square's position
	s.Lock()
	defer s.Unlock()
	s.Position.X += v.X * s.Speed * deltaTime
	s.Position.Y += v.Y * s.Speed * deltaTime
}

// Shoot creates new square's shot.
func (s *Square) Shoot(towards geometry.Point) {
	s.Lock()
	defer s.Unlock()

	// check if player can shoot
	if !s.CanShoot {
		return
	}

	// go through every slot for bullet
	for i := 0; i < len(s.Bullets); i++ {
		// if there is an empty slot
		if s.Bullets[i] == nil {
			// create a new bullet
			s.Bullets[i] = s.createBullet(towards)

			// disable shooting
			s.CanShoot = false

			// disable invulnerability
			select {
			case s.ShotCh <- struct{}{}:
			default:
			}

			// start weapon reloading
			go func() {
				s.reloadWeapon()
			}()

			return
		}
	}
}

// createBullet creates new bullet and adds it
// to Square's bullets.
//
// Accepts a point to shoot towards.
//
// Returns a pointer to the created Bullet object.
func (s *Square) createBullet(towards geometry.Point) *Bullet {
	// create and init bullet
	squareHalf := s.Size / 2
	bullet := &Bullet{
		Position: geometry.Point{
			X: s.Position.X + squareHalf,
			Y: s.Position.Y + squareHalf,
		},
		Vector: s.getShotVector(towards),
		Size:   config.BulletSize(),
		Speed:  config.BulletSpeed(),
		Damage: config.BulletDamage(),
	}

	return bullet
}

// getShotVector creates normalized vector of a shot.
//
// Accepts a point to shoot towards.
//
// Returns the created vector.
func (s *Square) getShotVector(towards geometry.Point) geometry.Vector {
	// count shot direction
	vector := geometry.Vector{
		X: towards.X - s.Position.X,
		Y: towards.Y - s.Position.Y,
	}

	// normalize vector
	vector.Normalize()

	return vector
}

// reloadWeapon sets the timer for square weapon reloading.
func (s *Square) reloadWeapon() {
	// set a timer
	timer := time.NewTimer(msToReload * time.Millisecond)
	defer timer.Stop()

	// wait till the timer ends
	<-timer.C

	// set the shoot ability to true
	s.Lock()
	defer s.Unlock()
	s.CanShoot = true
}

// UpdateBullets updates all existing Square's bullets'
// positions.
//
// Accepts a delta time value to correct bullet's speed.
func (s *Square) UpdateBullets(deltaTime float32) {
	// go through every bullet
	for _, b := range s.Bullets {
		// if the bullet exists
		if b != nil {
			// update it
			b.UpdateBulletPosition(deltaTime)
		}
	}
}

// RemoveBullet removes bullet from the Square's
// bullets.
//
// Accepts the index of the bullet.
func (s *Square) RemoveBullet(index int) {
	s.Bullets[index] = nil
}

// GetDamage reduces the health
// and the speed of the Square that was shot.
//
// Accepts a pointer to the bullet that damaged the Square
// and a pointer to the Square that shot.
func (s *Square) GetDamage(b *Bullet, shooter *Square) {
	// reduce square's health
	s.Health -= b.Damage

	// check if square should die
	if s.Health <= 0 {
		shooter.Kills++
		s.Deaths++
		go func() {
			s.regenerate()
		}()

		return
	}

	// reduce square's speed
	s.Speed -= s.Speed / 10
}

// regenerate moves Square to the respawn point,
// updates health and stats data,
// starts Square's invulnerability timer.
func (s *Square) regenerate() {
	// update square's stats
	s.Lock()
	s.Health = config.SquareHealth()
	s.Speed = config.SquareSpeed()
	s.Size = config.SquareSize()

	// return the square to its spawn point
	s.Position = s.Spawn

	// set the square's vulnerability to false for some time
	s.Vulnerable = false

	// save the native color of the square
	nativeColor := s.Color
	s.Unlock()

	// create a new context to control the invulnerability time
	ctx, cancel := context.WithTimeout(context.Background(), invulnerabilitySeconds*time.Second)
	defer cancel()

	// create a ticker to change colors during invulnerability
	ticker := time.NewTicker(changeColorMs * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// set the square's vulnerability to true
			// set the native color to the square
			s.restoreVulnerability(nativeColor)
			return
		case <-ticker.C:
			// change the square's color
			s.Lock()
			s.Color = utils.RandomBrightColor()
			s.Unlock()
		case <-s.ShotCh:
			// if the square makes a shot
			// he loses his invulnerability
			s.restoreVulnerability(nativeColor)
			return
		}
	}
}

// restoreVulnerability sets the square's vulnerability
// to true and the color to the native one.
//
// Accepts a native rgba color to set.
func (s *Square) restoreVulnerability(c color.RGBA) {
	s.Lock()
	defer s.Unlock()
	s.Vulnerable = true
	// set the native color to the square
	s.Color = c
}
