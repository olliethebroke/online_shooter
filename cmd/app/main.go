package main

import (
	"online_shooter/internal/game"
	"online_shooter/internal/logger"
)

func main() {
	if err := game.Run(); err != nil {
		logger.Fatal("failed to run the app: ", err)
	}
}
