package config

import (
	"online_shooter/internal/logger"
	"online_shooter/internal/utils"
)

// SquareHealth returns a square's health from the config.
// If the square's health is not initialized method gets it
// from the environment.
//
// Returns the square's health value.
func SquareHealth() int32 {
	if config.SquareHealth == nil {
		// get the var from the environment
		squareHealth, err := utils.GetIntEnvVar(squareHealthEnvName)
		if err != nil {
			logger.Fatal(err)
		}

		// store square's health value in the config
		config.SquareHealth = &squareHealth
	}

	return *config.SquareHealth
}

// SquareSize returns a square's size from the config.
// If the square's size is not initialized method gets it
// from the environment.
//
// Returns the square's size value.
func SquareSize() float32 {
	if config.SquareSize == nil {
		// get the var from the environment
		squareSize, err := utils.GetFloatEnvVar(squareSizeEnvName)
		if err != nil {
			logger.Fatal(err)
		}

		// store square's size value in the config
		config.SquareSize = &squareSize
	}

	return *config.SquareSize
}

// SquareSpeed returns a square's speed from the config.
// If the square's speed is not initialized method gets it
// from the environment.
//
// Returns the square's speed value.
func SquareSpeed() float32 {
	if config.SquareSpeed == nil {
		// get the var from the environment
		squareSpeed, err := utils.GetFloatEnvVar(squareSpeedEnvName)
		if err != nil {
			logger.Fatal(err)
		}

		// store square's speed value in the config
		config.SquareSpeed = &squareSpeed
	}

	return *config.SquareSpeed
}
