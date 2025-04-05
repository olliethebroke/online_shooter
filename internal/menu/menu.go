package menu

import (
	"online_shooter/internal/config"
	"online_shooter/internal/game/arena"
	"time"
)

const (
	MainMenuState             = 0
	ServerSettingsMenuState   = 1
	ServerConnectionMenuState = 2
)

type Menu struct {
	State              int8
	ConnectToServerBtn *Button
	StartServerBtn     *Button
	ConnectionSettings
	ServerSettings
	Active         bool
	lastChangeTime time.Time
}

type ServerSettings struct {
	PlayerCount   int
	ObstacleLevel string
	IsPublic      bool
}

type ConnectionSettings struct {
	ConnectionAddress string
	IpInput           TextInput
	PortInput         TextInput
	ActiveInputField  *TextInput
}

// NewMenu creates and initializes
// a new menu instance.
//
// Returns a pointer to
// the created instance.
func NewMenu() *Menu {
	var buttonWidth = config.ScreenWidth() / 5
	var buttonHeight = config.ScreenHeight() / 15
	m := &Menu{
		State: MainMenuState,
		ConnectToServerBtn: &Button{
			X:      (config.ScreenWidth() - buttonWidth) / 2,
			Y:      config.ScreenHeight()/2 - buttonHeight,
			Width:  buttonWidth,
			Height: buttonHeight,
			Label:  "Connect to Server",
		},
		StartServerBtn: &Button{
			X:      (config.ScreenWidth() - buttonWidth) / 2,
			Y:      config.ScreenHeight()/2 + buttonHeight,
			Width:  buttonWidth,
			Height: buttonHeight,
			Label:  "Start Server",
		},
		ServerSettings: ServerSettings{
			PlayerCount:   4,
			ObstacleLevel: arena.MediumObstaclesAmount,
			IsPublic:      true,
		},
		ConnectionSettings: ConnectionSettings{
			IpInput: TextInput{
				Value:     "127.0.0.1",
				MaxLength: 15,
				IsActive:  true,
			},
			PortInput: TextInput{
				Value:     "8080",
				MaxLength: 5,
			},
		},
		Active: true,
	}
	m.ActiveInputField = &m.IpInput
	return m
}
