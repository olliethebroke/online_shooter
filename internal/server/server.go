package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"online_shooter/internal/game/game"
	"online_shooter/internal/logger"
	"online_shooter/internal/menu"
	"online_shooter/internal/model"
	"sync"
	"time"
)

const (
	publicInterface  = "0.0.0.0:8080"
	PrivateInterface = "localhost:8080"

	CreatePlayerPostfix  = "/player/create"
	ConnectPlayerPostfix = "/connect/"
)

type Server struct {
	game.Game
	serverMutex   sync.RWMutex
	playerUpdates map[int64]*model.PlayerUpdateMessage
	lastUpdate    *time.Time
}

// Run initializes and starts server listening an interface.
// Server accepts http requests and maintain websocket connection
// with clients using broadcast function to send the game state.
func (s *Server) Run(settings *menu.ServerSettings) {
	// create and init a new router instance
	r := chi.NewRouter()

	// set up the server
	s.setup(settings)

	// change tcp network address whether the server is public or not
	var url string
	if settings.IsPublic {
		url = publicInterface
	} else {
		url = PrivateInterface
	}

	// add handlers
	r.Post(CreatePlayerPostfix, s.createPlayerHandler)
	r.Get(ConnectPlayerPostfix+"{id}", s.connectPlayerHandler)

	// start broadcasting server state
	go s.broadcast()

	// listen on address
	err := http.ListenAndServe(url, r)
	if err != nil {
		logger.Fatal("error while listening on server: ", err)
	}
}

// setup initializes map fields and a new game of the server instance.
func (s *Server) setup(settings *menu.ServerSettings) {
	// init the map with updates
	s.playerUpdates = make(map[int64]*model.PlayerUpdateMessage)

	// inits a new game
	s.InitServerGame(settings)
}
