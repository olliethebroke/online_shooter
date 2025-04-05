package config

import (
	"online_shooter/internal/logger"
	"online_shooter/internal/utils"
)

// ScreenWidth returns a screen width from the config.
// If the screen width is not initialized method gets it
// from the environment.
//
// Returns the screen width value.
func ScreenWidth() float32 {
	if config.ScreenWidth == nil {
		// get the var from the environment
		screenWidth, err := utils.GetFloatEnvVar(screenWidthEnvName)
		if err != nil {
			logger.Fatal(err)
		}

		// store screen width value in the config
		config.ScreenWidth = &screenWidth
	}

	return *config.ScreenWidth
}

// ScreenHeight returns a screen height from the config.
// If the screen height is not initialized method gets it
// from the environment.
//
// Returns the screen height value.
func ScreenHeight() float32 {
	if config.ScreenHeight == nil {
		// get the var from the environment
		screenHeight, err := utils.GetFloatEnvVar(screenHeightEnvName)
		if err != nil {
			logger.Fatal(err)
		}

		// store screen height value in the config
		config.ScreenHeight = &screenHeight
	}

	return *config.ScreenHeight
}
