package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"net/http"
	"online_shooter/internal/game/entity"
	"online_shooter/internal/logger"
	"online_shooter/internal/model"
	"online_shooter/internal/utils"
)

// upgrader is used for creating a websocket connection from http connection
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// createPlayerHandler handles an http request from the client
// creating a new player in the game and response sending
// an arena and the created player instances to the client.
func (s *Server) createPlayerHandler(w http.ResponseWriter, r *http.Request) {
	// create and init a new player instance
	player := entity.NewPlayer(nil)

	// add the created player to the server as a new player
	s.addPlayer(player)

	// send created status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// encode a response providing the created player and the game arena
	response := &model.CreatePlayerResponse{
		Arena:  s.Arena,
		Player: player,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "failed to encode the response", http.StatusInternalServerError)
		logger.Warn("error while encoding create player response: ", err)
		return
	}
}

// connectPlayerHandler handles an http request from the client
// establishing websocket connection between the client and the server.
// As a result of the method starts a readMessage method that
// provides eternal listening of client updates.
func (s *Server) connectPlayerHandler(w http.ResponseWriter, r *http.Request) {
	// get player id from the url param
	id, err := utils.StringToInt64(chi.URLParam(r, "id"))
	if err != nil {
		logger.Warn("error while parsing url param: ", err)
		return
	}

	// create and init a new websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Warn("error while upgrading connection: ", err)
		return
	}

	// update player's connection field
	s.serverMutex.Lock()
	s.Squares[id].Conn = conn
	player := s.Squares[id]
	s.serverMutex.Unlock()

	// create a callback function to remove the player from the server when they disconnect
	// this function will be called by ReadMessages if the player's connection is closed
	removePlayerFunc := func(id int64) {
		conn.Close()
		s.removePlayer(id)
	}

	// log the player connection
	logger.Info(fmt.Sprintf("player%d connected to the server", id))

	// start reading messages from the player
	// the callback function is passed to handle player disconnection
	go s.readMessages(player, removePlayerFunc)
}
