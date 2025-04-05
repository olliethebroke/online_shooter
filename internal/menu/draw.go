package menu

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"image/color"
	"online_shooter/internal/assets"
	"time"
)

// Draw draws the application's menu.
//
// Accepts a pointer to the image as an argument.
func (m *Menu) Draw(screen *ebiten.Image) {
	// count a place of the menu module header
	headerY := screen.Bounds().Dy() / 10

	// draw the module
	switch m.State {
	// if it is a main module
	case MainMenuState:
		m.drawMainMenu(screen, headerY)

	// if it is a server settings module
	case ServerSettingsMenuState:
		m.drawServerSettingsMenu(screen, headerY)

	// if it is a server connection module
	case ServerConnectionMenuState:
		m.drawConnectionSettingsMenu(screen, headerY)
	}
}

// drawMainMenu draws the main menu module.
//
// Accepts a pointer to the image object and a y
// coordinate of the header as arguments.
func (m *Menu) drawMainMenu(screen *ebiten.Image, headerY int) {
	// draw the header
	drawCenteredText(screen, "menu", headerY, color.White)

	// draw the "Connect to Server" button
	drawButton(m.ConnectToServerBtn, screen)

	// draw the "Start Server" button
	drawButton(m.StartServerBtn, screen)
}

// drawServerSettingsMenu draws the server settings menu module.
//
// Accepts a pointer to the image object and a y
// coordinate of the header as arguments.
func (m *Menu) drawServerSettingsMenu(screen *ebiten.Image, headerY int) {
	// draw the header
	drawCenteredText(screen, "Server Settings", headerY, color.White)

	// draw parameters
	drawCenteredText(screen, fmt.Sprintf("Players Amount : %d", m.PlayerCount), headerY*2, color.White)
	drawCenteredText(screen, fmt.Sprintf("Obstacles Amount: %s", m.ObstacleLevel), headerY*3, color.White)
	drawCenteredText(screen, fmt.Sprintf("Is Server Public: %v", m.IsPublic), headerY*4, color.White)

	// draw the "Start Server" button
	drawButton(m.StartServerBtn, screen)

	// draw the hint
	hintY := float32(screen.Bounds().Dy()) * 0.9
	drawCenteredText(screen, "Use Arrow Keys and Space to Change Settings", int(hintY), color.White)
}

// drawConnectionSettingsMenu draws the connection settings menu module.
//
// Accepts a pointer to the image object and a y
// coordinate of the header as arguments.
func (m *Menu) drawConnectionSettingsMenu(screen *ebiten.Image, headerY int) {
	// draw the header
	drawCenteredText(screen, "Connection Settings", headerY, color.White)

	// draw ip input
	drawInputField(screen, &m.IpInput, "Server IP: ", headerY*2)

	// draw port input
	drawInputField(screen, &m.PortInput, "Server Port: ", headerY*3)

	// draw the "Start Server" button
	drawButton(m.ConnectToServerBtn, screen)

	// draw the hint
	hintY := float32(screen.Bounds().Dy()) * 0.9
	drawCenteredText(screen, "Tab - switch field, Enter - confirm", int(hintY), color.White)
}

// drawCenteredText draws a text in the center of the screen.
//
// Accepts a pointer to the screen, a string that needs to be printed,
// a y coord of the text, and the color of the text.
func drawCenteredText(screen *ebiten.Image, s string, y int, clr color.Color) {
	// get text width
	textWidth := font.MeasureString(assets.Font(), s)

	// text width to pixels
	textWidthPx := textWidth.Ceil()

	// count x coord
	x := (screen.Bounds().Dx() - textWidthPx) / 2

	// get text height
	ascent := assets.Font().Metrics().Ascent.Ceil()

	// draw text
	text.Draw(screen, s, assets.Font(), x, y+ascent, clr)
}

// drawInputField draws an input field on the screen.
// Accepts a pointer to the image object, a pointer to the
// text input object that needs to be drawn, its label and an y coord
// where to place field on the screen as arguments.
func drawInputField(screen *ebiten.Image, input *TextInput, label string, y int) {
	// get text width
	textWidth := font.MeasureString(assets.Font(), label)

	// text width to pixels
	textWidthPx := textWidth.Ceil()

	// draw label
	text.Draw(screen, label, assets.Font(), (screen.Bounds().Dx()-textWidthPx)/2, y, color.White)

	// draw the input field
	value := input.Value
	if input.IsActive {
		// add cursor blinking animation
		if time.Now().UnixNano()/1e8%2 == 0 {
			value = value[:input.CursorPos] + "|" + value[input.CursorPos:]
		}
	}
	text.Draw(screen, value, assets.Font(), (screen.Bounds().Dx()+textWidthPx)/2, y, color.White)
}

// drawButton draws the button on the screen.
//
// Accepts a pointer to the button to draw
// and screen objects as arguments.
func drawButton(btn *Button, screen *ebiten.Image) {
	// draw the button
	vector.DrawFilledRect(
		screen,
		btn.X,
		btn.Y,
		btn.Width,
		btn.Height,
		color.RGBA{R: 100, G: 100, B: 100, A: 255},
		true,
	)

	// draw the button text
	drawCenteredText(screen, btn.Label, int(btn.Y+btn.Height/2-float32(assets.Font().Metrics().Ascent.Ceil())), color.White)
}
