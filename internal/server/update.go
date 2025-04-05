package server

import (
	"online_shooter/internal/game/entity"
	"online_shooter/internal/game/geometry"
	"online_shooter/internal/logger"
	"online_shooter/internal/model"
	"time"
)

const framesPerSecond = 60
const refreshingRate = 1000 / framesPerSecond

// Update updates a server's game state using
// the data about the players from the clients.
func (s *Server) Update() {
	now := time.Now()

	// set the last update time if it doesn't exist
	if s.lastUpdate == nil {
		s.lastUpdate = &now
		return
	}

	// count the speed correction value
	deltaTime := float32(now.Sub(*s.lastUpdate).Seconds())
	s.lastUpdate = &now

	// go through every square in the game and update its state
	for _, square := range s.Squares {
		if square.IsBot {
			// change game's state for the bot
			enemy, distance := s.FindEnemy(square)
			square.Move(s.CountMovingVector(square, enemy, distance), deltaTime)
			aim := s.CountShootingPoint(enemy, distance)
			if aim != nil {
				square.Shoot(*aim)
			}
		} else {
			// change game's state for the player
			if s.playerUpdates[square.Id] != nil {
				updatePlayer(square, s.playerUpdates[square.Id], deltaTime)
				s.playerUpdates[square.Id] = nil
			}
		}

		s.CheckSquareCollision(square)
		square.UpdateBullets(deltaTime)
		s.CheckBulletsCollision(square)
	}
}

// updatePlayer updates the player's square state
// using the information from the client.
//
// Accepts a pointer to the player instance,
// a pointer to the instance with a square's state update
// and a delta time value to correct the player's square speed.
func updatePlayer(player *entity.Square, upd *model.PlayerUpdateMessage, deltaTime float32) {
	// change player's position
	vector := &geometry.Vector{}
	if upd.UpKeyPressed {
		vector.Y--
	}
	if upd.DownKeyPressed {
		vector.Y++
	}
	if upd.LeftKeyPressed {
		vector.X--
	}
	if upd.RightKeyPressed {
		vector.X++
	}

	vector.Normalize()
	player.Move(*vector, deltaTime)

	// make player shoot
	if upd.Shot {
		player.Shoot(upd.Aim)
	} else {
		player.Lock()
		player.CanShoot = true
		player.Unlock()
	}
}

// addPlayer adds a new player to the game swapping
// it with the bot. If there are no bots in the game
// logs it and doesn't add the player.
//
// Accepts a pointer to the player that should be added.
func (s *Server) addPlayer(player *entity.Square) {
	s.serverMutex.Lock()

	// find the weakest bot to remove it from the game
	bot := s.FindWeakestBot()
	if bot != nil {
		// if there is a bot in the game
		// generate an id
		player.Id = s.GenerateUniqueId()

		// transfer the bot's spawn point to the player
		player.Spawn = bot.Spawn

		// set the player's position to its spawn point
		player.Position = player.Spawn

		// delete the bot
		close(s.Squares[bot.Id].ShotCh)
		delete(s.Squares, bot.Id)

		// add the player to the game
		s.Squares[player.Id] = player
	} else {
		// if there is no bot in the game
		// log it
		logger.Info("there is no space for a new player")

		// close connection with the player that wanted to connect
		player.Conn.Close()
	}

	s.serverMutex.Unlock()
}

// removePlayer removes a player with accepted id
// from the game.
//
// Accepts an id of the player that should be removed.
func (s *Server) removePlayer(id int64) {
	s.serverMutex.Lock()
	// close player connection
	s.Squares[id].Conn.Close()

	// save player's spawn point
	spawn := s.Squares[id].Spawn

	// delete player
	close(s.Squares[id].ShotCh)
	delete(s.Squares, id)

	// generate new bot id
	id = s.GenerateUniqueId()

	// create and init a new bot instance
	bot := entity.NewBot(id)

	// transfer the deleted player's spawn point to the bot
	bot.Spawn = spawn

	// add the created bot to the game
	s.Squares[id] = bot
	s.serverMutex.Unlock()
}
