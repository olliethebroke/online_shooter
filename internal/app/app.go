package app

import (
	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"online_shooter/internal/config"
	"online_shooter/internal/game/game"
	"online_shooter/internal/menu"
	"online_shooter/internal/server"
)

// init initializes the application.
func init() {
	// load variables from the env file
	config.Load("./app.env")
}

type App struct {
	screenWidth  float32
	screenHeight float32
	game         *game.Game
	menu         *menu.Menu
	server       *server.Server
	conn         *websocket.Conn
}

func Run() error {
	app := &App{
		screenWidth:  config.ScreenWidth(),
		screenHeight: config.ScreenHeight(),
		game:         &game.Game{},
		menu:         menu.NewMenu(),
	}
	ebiten.SetWindowSize(int(app.screenWidth), int(app.screenHeight))
	ebiten.SetWindowTitle("Shooter")
	err := ebiten.RunGame(app)
	return err
}
