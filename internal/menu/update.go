package menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"online_shooter/internal/event"
	"online_shooter/internal/game/arena"
	"time"
)

const waitTillChangeInMs = 150

// Update updates menu state in case it is active.
//
// Returns event that must be done after user's
// interaction with menu.
func (m *Menu) Update() event.Event {
	// read user interaction with keyboard and mouse
	// change menu parameters
	return m.readUserInteraction()
}

// readUserInteraction checks if user interacts
// with a mouse and a keyboard to change application's state.
func (m *Menu) readUserInteraction() event.Event {
	switch m.State {
	// if it is the main menu state
	case MainMenuState:
		// check if user interacts with mouse to click on the buttons
		m.readMainMenuInteraction()

	// if it is the settings menu state
	case ServerSettingsMenuState:
		// check if user interacts with mouse
		// and keyboard to change parameters and click the button
		return m.readSettingMenuInteraction()

	// if it is a connection menu state
	case ServerConnectionMenuState:
		// check if user interacts with mouse
		// and keyboard to change connection address
		return m.readConnectionMenuInteraction()
	}

	return -1
}

// readMainMenuInteraction checks if user interacts with a mouse
// to click on the buttons in the main menu.
func (m *Menu) readMainMenuInteraction() {
	// check if lbm is pressed
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		// check Connect to Server button is clicked
		if m.ConnectToServerBtn.IsClicked(float32(x), float32(y)) {
			m.State = ServerConnectionMenuState
		}

		// check Start Server button is clicked
		if m.StartServerBtn.IsClicked(float32(x), float32(y)) {
			m.State = ServerSettingsMenuState
		}
		m.lastChangeTime = time.Now()
	}
}

// readSettingMenuInteraction checks if user interacts
// with a mouse and a keyboard to change the parameters
// and click on the button in the settings menu.
func (m *Menu) readSettingMenuInteraction() event.Event {
	// prevent too fast changing of the parameter
	now := time.Now()
	if now.Sub(m.lastChangeTime) < waitTillChangeInMs*time.Millisecond {
		return -1
	}

	// if arrow up key is pressed
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		// increase players amount
		m.PlayerCount++
		if m.PlayerCount > 100 {
			m.PlayerCount = 100
		}
		m.lastChangeTime = now
	}
	// if arrow down key is pressed
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		// decrease players amount
		m.PlayerCount--
		if m.PlayerCount < 2 {
			m.PlayerCount = 2
		}
		m.lastChangeTime = now
	}
	// if arrow right key is pressed
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		// switch obstacles amount level
		if m.ObstacleLevel == arena.LowObstaclesAmount {
			m.ObstacleLevel = arena.MediumObstaclesAmount
		} else if m.ObstacleLevel == arena.MediumObstaclesAmount {
			m.ObstacleLevel = arena.HighObstaclesAmount
		} else if m.ObstacleLevel == arena.HighObstaclesAmount {
			m.ObstacleLevel = arena.LowObstaclesAmount
		}
		m.lastChangeTime = now
	}
	// if arrow left key is pressed
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		// switch obstacles amount level
		if m.ObstacleLevel == arena.LowObstaclesAmount {
			m.ObstacleLevel = arena.HighObstaclesAmount
		} else if m.ObstacleLevel == arena.MediumObstaclesAmount {
			m.ObstacleLevel = arena.LowObstaclesAmount
		} else if m.ObstacleLevel == arena.HighObstaclesAmount {
			m.ObstacleLevel = arena.MediumObstaclesAmount
		}
		m.lastChangeTime = now
	}

	// if space is pressed
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		// switch server's access mode
		m.IsPublic = !m.IsPublic
		m.lastChangeTime = now
	}

	// check if lbm is pressed
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		// check Start Server button is clicked
		if m.StartServerBtn.IsClicked(float32(x), float32(y)) {
			// close menu
			m.Active = false

			// start server
			return event.EventStartServer
		}
	}

	return -1
}

// readConnectionMenuInteraction checks if user interacts
// with a mouse and a keyboard to change the address
// and click on the button in the connection menu.
func (m *Menu) readConnectionMenuInteraction() event.Event {
	// prevent too fast changing of the updates
	now := time.Now()
	if now.Sub(m.lastChangeTime) < waitTillChangeInMs*time.Millisecond {
		return -1
	}

	// field switching logic
	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		if m.ActiveInputField == &m.IpInput {
			m.IpInput.IsActive = false
			m.PortInput.IsActive = true
			m.ActiveInputField = &m.PortInput
		} else {
			m.IpInput.IsActive = true
			m.PortInput.IsActive = false
			m.ActiveInputField = &m.IpInput
		}
		m.lastChangeTime = now
	}

	// text input in the active field
	if m.ActiveInputField != nil {
		// add symbols
		for _, r := range ebiten.AppendInputChars(nil) {
			if m.ActiveInputField.MaxLength == 0 || len(m.ActiveInputField.Value) < m.ActiveInputField.MaxLength {
				m.ActiveInputField.Value = m.ActiveInputField.Value[:m.ActiveInputField.CursorPos] + string(r) +
					m.ActiveInputField.Value[m.ActiveInputField.CursorPos:]
				m.ActiveInputField.CursorPos++
				m.lastChangeTime = now
			}
		}

		// delete symbol
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && m.ActiveInputField.CursorPos > 0 {
			m.ActiveInputField.Value = m.ActiveInputField.Value[:m.ActiveInputField.CursorPos-1] +
				m.ActiveInputField.Value[m.ActiveInputField.CursorPos:]
			m.ActiveInputField.CursorPos--
			m.lastChangeTime = now
		}

		// move cursor
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && m.ActiveInputField.CursorPos > 0 {
			m.ActiveInputField.CursorPos--
			m.lastChangeTime = now
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyRight) && m.ActiveInputField.CursorPos < len(m.ActiveInputField.Value) {
			m.ActiveInputField.CursorPos++
			m.lastChangeTime = now
		}
	}

	// check if the lbm is pressed
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// get click coords
		clickX, clickY := ebiten.CursorPosition()

		// check if the mouse covered a button during the click
		if m.ConnectToServerBtn.IsClicked(float32(clickX), float32(clickY)) {
			// close menu
			m.Active = false

			// build the connection address
			// and return connect to server event
			m.ConnectionAddress = m.IpInput.Value + ":" + m.PortInput.Value
			return event.EventConnectToServer
		}

	}

	return -1
}
