package main

import (
	"online_shooter/internal/app"
	"online_shooter/internal/logger"
)

func main() {
	// start the application
	if err := app.Run(); err != nil {
		logger.Fatal("failed to run the game: ", err)
	}
}
