package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"online_shooter/internal/game/entity"
	"online_shooter/internal/logger"
	"online_shooter/internal/model"
	"time"
)

// broadcast provides every connected client with
// actual game state using WebSocket connection.
func (s *Server) broadcast() {
	// set refreshing time for the ticker
	ticker := time.NewTicker(refreshingRate * time.Millisecond)
	defer ticker.Stop()

	// every tick updates a game state and broadcasts data to clients
	for range ticker.C {
		s.serverMutex.Lock()

		// update game state
		s.Update()

		// create and init a new update instance
		s.Arena.ArenaMutex.RLock()
		s.Game.GameMutex.RLock()
		gameUpdate := &model.GameUpdateMessage{
			Obstacles: s.Arena.Obstacles,
			Squares:   s.Squares,
		}

		// lock every obstacle before marshalling
		for _, obstacle := range gameUpdate.Obstacles {
			obstacle.Lock()
		}
		// lock every square before marshalling
		for _, square := range gameUpdate.Squares {
			square.Lock()
		}

		// marshal the update
		msg, _ := json.Marshal(gameUpdate)
		// unlock every obstacle before marshalling
		for _, obstacle := range gameUpdate.Obstacles {
			obstacle.Unlock()
		}
		// unlock every square before marshalling
		for _, square := range gameUpdate.Squares {
			square.Unlock()
		}
		s.Game.GameMutex.RUnlock()
		s.Arena.ArenaMutex.RUnlock()

		// send the update for every player
		for _, square := range s.Squares {
			if !square.IsBot {
				square.RLock()
				if err := square.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					logger.Warn("error while broadcasting game state: ", err)
				}
				square.RUnlock()
			}
		}

		s.serverMutex.Unlock()
	}
}

// readMessages reads messages from the client about the player
// game state in infinite loop and updates the player state
// on the server side. if the connection between client and server
// is closed the player is removed from the game.
//
// Accepts a function to remove the player as an argument.
func (s *Server) readMessages(player *entity.Square, remove func(id int64)) {
	// run an infinite loop reading messages from the client
	for {
		// read the message
		player.RLock()
		_, msg, err := player.Conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				logger.Info("player disconnected: ", player.Id)
			} else {
				logger.Warn("failed to read a message from the client ", player.Id, ": ", err)

			}
			break
		}
		player.RUnlock()

		// decode the message
		var updatePlayerMessage model.PlayerUpdateMessage
		err = json.Unmarshal(msg, &updatePlayerMessage)
		if err != nil {
			logger.Warn("failed to decode a message from the client ", player.Id, ": ", err)
		} else {
			// add player update to the slice
			s.serverMutex.Lock()
			s.playerUpdates[player.Id] = &updatePlayerMessage
			s.serverMutex.Unlock()
		}
	}

	// remove the player from the game
	remove(player.Id)
}
