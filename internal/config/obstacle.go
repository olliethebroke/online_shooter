package config

import (
	"online_shooter/internal/logger"
	"online_shooter/internal/utils"
)

// ObstacleHealth returns an obstacle's health from the config.
// If the obstacle's health is not initialized method gets it
// from the environment.
//
// Returns the obstacle's health value.
func ObstacleHealth() int32 {
	if config.ObstacleHealth == nil {
		// get the var from the environment
		obstacleHealth, err := utils.GetIntEnvVar(obstacleHealthEnvName)
		if err != nil {
			logger.Fatal(err)
		}

		config.ObstacleHealth = &obstacleHealth

	}

	return *config.ObstacleHealth
}

// ObstacleSize returns an obstacle's size from the config.
// If the obstacle's size is not initialized method gets it
// from the environment.
//
// Returns the obstacle's size value.
func ObstacleSize() float32 {
	if config.ObstacleSize == nil {
		// get the var from the environment
		obstacleSize, err := utils.GetFloatEnvVar(obstacleSizeEnvName)
		if err != nil {
			logger.Fatal(err)
		}

		// store obstacle's size value in the config
		config.ObstacleSize = &obstacleSize
	}

	return *config.ObstacleSize
}
