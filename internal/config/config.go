package config

import (
	"github.com/joho/godotenv"
	"online_shooter/internal/logger"
)

const (
	screenWidthEnvName  = "SCREEN_WIDTH"
	screenHeightEnvName = "SCREEN_HEIGHT"

	squareHealthEnvName = "SQUARE_HEALTH"
	squareSizeEnvName   = "SQUARE_SIZE"
	squareSpeedEnvName  = "SQUARE_SPEED"

	obstacleHealthEnvName = "OBSTACLE_HEALTH"
	obstacleSizeEnvName   = "OBSTACLE_SIZE"

	bulletDamageEnvName = "BULLET_DAMAGE"
	bulletSizeEnvName   = "BULLET_SIZE"
	bulletSpeedEnvName  = "BULLET_SPEED"
)

type GameConfig struct {
	ScreenWidth  *float32
	ScreenHeight *float32

	SquareHealth *int32
	SquareSize   *float32
	SquareSpeed  *float32

	ObstacleHealth *int32
	ObstacleSize   *float32

	BulletDamage *int32
	BulletSize   *float32
	BulletSpeed  *float32
}

var config = GameConfig{}

// Load loads variables from env file
// to the process environmental variables.
//
// Accepts path to the env file.
func Load(path string) {
	if err := godotenv.Load(path); err != nil {
		logger.Fatal("failed to load variables from env file: ", err)
	}
}
