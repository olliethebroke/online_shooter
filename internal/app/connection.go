package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"online_shooter/internal/logger"
	"online_shooter/internal/model"
	"online_shooter/internal/server"
)

// connectWithServer creates a new player on the server and
// establishes WebSocket connection with the server to communicate
// with it.
func (a *App) connectWithServer() {
	// create a new player
	err := a.createPlayer()
	if err != nil {
		logger.Warn("error while creating player on server: ", err)
		return
	}

	// get a websocket connection with the server
	err = a.setupWebSocketConnection()
	if err != nil {
		logger.Warn("error while creating websocket connection with server: ", err)
		return
	}
}

// createPlayer creates a new player
// on the server and gets its id.
//
// Returns an error if the creation fails.
func (a *App) createPlayer() error {
	// make the request
	url := fmt.Sprintf("http://%s%s",
		a.menu.ConnectionAddress,
		server.CreatePlayerPostfix)
	resp, err := http.Post(url, "", nil)
	if err != nil {
		return err
	}

	// decode a createPlayerResponse
	var createPlayerResponse model.CreatePlayerResponse
	err = json.NewDecoder(resp.Body).Decode(&createPlayerResponse)
	if err != nil {
		return err
	}

	// set the game fields in order to the server
	a.game.Player = createPlayerResponse.Player
	a.game.Arena = createPlayerResponse.Arena

	return nil
}

// setupWebSocketConnection establishes a WebSocket connection to the server.
// It uses the IP, port, and connection endpoint provided in the app's menu configuration.
//
// Returns an error if the connection fails.
func (a *App) setupWebSocketConnection() error {
	// create a WebSocket dialer to initiate the connection
	dialer := websocket.Dialer{}

	// construct the WebSocket URL using the IP, port, and connection endpoint
	endpoint := fmt.Sprintf("%s%d", server.ConnectPlayerPostfix, a.game.Player.Id)
	wsURL := fmt.Sprintf("ws://%s%s",
		a.menu.ConnectionAddress,
		endpoint)

	// dial the server to establish a WebSocket connection
	conn, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		return err
	}

	// store the WebSocket connection in the app for future use
	a.conn = conn

	return nil
}
