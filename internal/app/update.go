package app

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"online_shooter/internal/event"
	"online_shooter/internal/game/game"
	"online_shooter/internal/logger"
	"online_shooter/internal/model"
	"online_shooter/internal/server"
	"time"
)

// Update updates the application by one tick.
//
// Update updates only the game logic and Draw draws the screen.
func (a *App) Update() error {
	// update the menu if it is required
	if a.menu.Active {
		// get an event
		e := a.menu.Update()

		// check the menu interaction event
		switch e {
		// if it is a Connect to Server event
		case event.EventConnectToServer:
			// run the client game
			a.runGame()

		// if it is a Start Server event
		case event.EventStartServer:
			// init the server
			a.server = &server.Server{}

			// start the server
			go a.server.Run(&a.menu.ServerSettings)

			// set the connection address
			a.menu.ConnectionAddress = server.PrivateInterface

			// run the client game
			a.runGame()
		}
	}

	// update the game if it is required
	if a.game.Active {
		a.game.GameMutex.RLock()

		// update server in case keys are pressed
		movement := a.game.Player.GetPlayerMovement()

		// update server in case lbm is pressed
		shooting := a.game.Player.GetPlayerShooting(a.game.Camera)

		// create and init a pointer to the PlayerUpdateMessage instance
		playerUpdate := &model.PlayerUpdateMessage{
			UpKeyPressed:    movement.UpKeyPressed,
			DownKeyPressed:  movement.DownKeyPressed,
			LeftKeyPressed:  movement.LeftKeyPressed,
			RightKeyPressed: movement.RightKeyPressed,
			Shot:            shooting.Shot,
			Aim:             shooting.Aim,
		}

		a.game.GameMutex.RUnlock()
		// send the update to the server
		a.sendUpdateToServer(playerUpdate)
	}

	return nil
}

func (a *App) runGame() {
	// try to connect to the server
	for {
		// connect to the server
		a.connectWithServer()
		if a.conn != nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	// init the client game
	a.game.InitClientGame()

	// start reading updates from the server
	go a.readServerUpdates()
}

// readServerUpdates reads game state updates from the server
// using websocket connection and refreshes the client game state.
func (a *App) readServerUpdates() {
	defer a.conn.Close()

	// run an infinite loop for reading messages from the server
	for {
		// read a message from the server
		_, msg, err := a.conn.ReadMessage()
		if err != nil || !a.game.Active {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				logger.Info("server disconnected")
			} else {
				logger.Warn("failed to read a message from the server: ", err)
			}
			break
		}

		// decode the message from the server
		var gameUpdate model.GameUpdateMessage
		err = json.Unmarshal(msg, &gameUpdate)
		if err != nil {
			logger.Warn("failed to decode a message from the server: ", err)
		}

		// update client game
		updateGame(a.game, &gameUpdate)
	}
}

// updateGame updates a client's game state using
// the received data about the game.
//
// Accepts a pointer to the update data.
func updateGame(g *game.Game, gameUpdate *model.GameUpdateMessage) {
	g.GameMutex.Lock()
	defer g.GameMutex.Unlock()

	// update game squares on the client using data from the server
	g.Squares = gameUpdate.Squares

	// update user's square
	g.Player = g.Squares[g.Player.Id]

	// update game obstacles on the client side using data from the server
	g.Arena.Obstacles = gameUpdate.Obstacles

	// move the game Camera to the player
	if g.Player != nil {
		g.Camera.Move(g.Player.Position)
	}
}

// sendUpdateToServer writes an update player message to the server
// using WebSocket connection.
//
// Accepts a pointer to the PlayerUpdateMessage instance as an argument.
func (a *App) sendUpdateToServer(playerUpdate *model.PlayerUpdateMessage) {
	msg, _ := json.Marshal(playerUpdate)
	if err := a.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		logger.Warn("failed to send update message to the server: ", err)
	}
}
