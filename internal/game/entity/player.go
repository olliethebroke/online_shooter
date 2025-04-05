package entity

import (
	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"online_shooter/internal/config"
	"online_shooter/internal/game/camera"
	"online_shooter/internal/game/geometry"
	"online_shooter/internal/utils"
)

type Movement struct {
	LeftKeyPressed  bool
	UpKeyPressed    bool
	RightKeyPressed bool
	DownKeyPressed  bool
}

type Shooting struct {
	Shot bool
	Aim  geometry.Point
}

// NewPlayer creates and initializes
// new player square instance with default parameters.
//
// Accepts a pointer to the websocket connection
// as an argument.
//
// Returns pointer to the created player square.
func NewPlayer(conn *websocket.Conn) *Square {
	p := &Square{
		Conn:       conn,
		Health:     100,
		Speed:      config.SquareSpeed(),
		Size:       config.SquareSize(),
		Color:      utils.RandomBrightColor(),
		CanShoot:   false,
		Vulnerable: true,
		ShotCh:     make(chan struct{}),
	}
	return p
}

// GetPlayerMovement registers player's moving in case
// keys are pressed.
//
// Returns a pointer to a struct containing all player movements.
func (s *Square) GetPlayerMovement() *Movement {
	movement := &Movement{}

	// change update player data depending on pressed keys
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		movement.UpKeyPressed = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		movement.DownKeyPressed = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		movement.RightKeyPressed = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		movement.LeftKeyPressed = true
	}

	return movement
}

// GetPlayerShooting gets the point to shoot toward
// and tries to make a shot in case
// LBM is pressed.
//
// Accepts a pointer to the camera object to recount
// the mouse coordinates due to the world coordinate system.
//
// Returns a pointer to a struct containing info about a player's shot.
func (s *Square) GetPlayerShooting(camera *camera.Camera) *Shooting {
	shooting := &Shooting{}

	// check if lbm is pressed
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// get cursor position
		x, y := ebiten.CursorPosition()
		towards := geometry.Point{
			X: float32(x),
			Y: float32(y),
		}

		// recount the point to shoot
		// to the world coordinate system
		towards = camera.ScreenToWorld(towards)

		// set the shot flog to true for server update
		shooting.Shot = true

		// set the aim for server update
		shooting.Aim = towards
	}

	return shooting
}

// CountBulletsAmount returns an amount of bullets
// which are available for the player to use.
func (s *Square) CountBulletsAmount() uint8 {
	var counter uint8

	// go through every bullet pointer
	for _, b := range s.Bullets {
		// and check if it is null - bullet unused
		if b == nil {
			// increment counter
			counter++
		}
	}

	return counter
}
