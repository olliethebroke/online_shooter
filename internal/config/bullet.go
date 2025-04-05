package config

import (
	"online_shooter/internal/logger"
	"online_shooter/internal/utils"
)

// BulletDamage returns a bullet's damage from the config.
// If the bullet's damage is not initialized method gets it
// from the environment.
//
// Returns the bullet's damage value.
func BulletDamage() int32 {
	if config.BulletDamage == nil {
		// get the var from the environment
		bulletDamage, err := utils.GetIntEnvVar(bulletDamageEnvName)
		if err != nil {
			logger.Fatal(err)
		}

		// store bullet's damage value in the config
		config.BulletDamage = &bulletDamage
	}

	return *config.BulletDamage
}

// BulletSize returns a bullet's size from the config.
// If the bullet's size is not initialized method gets it
// from the environment.
//
// Returns the bullet's size value.
func BulletSize() float32 {
	if config.BulletSize == nil {
		// get the var from the environment
		bulletSize, err := utils.GetFloatEnvVar(bulletSizeEnvName)
		if err != nil {
			logger.Fatal(err)
		}

		// store bullet's size value in the config
		config.BulletSize = &bulletSize
	}

	return *config.BulletSize
}

// BulletSpeed returns a bullet's speed from the config.
// If the bullet's speed is not initialized method gets it
// from the environment.
//
// Returns the bullet's speed value.
func BulletSpeed() float32 {
	if config.BulletSpeed == nil {
		// get the var from the environment
		bulletSpeed, err := utils.GetFloatEnvVar(bulletSpeedEnvName)
		if err != nil {
			logger.Fatal(err)
		}

		// store square's speed value in the config
		config.BulletSpeed = &bulletSpeed
	}

	return *config.BulletSpeed
}
