package menu

type Button struct {
	X, Y, Width, Height float32
	Label               string
}

// IsClicked checks if the button is clicked.
//
// Accepts coords of the click.
//
// Returns the bool value answering
// "Is the button clicked?".
func (b *Button) IsClicked(x, y float32) bool {
	return x >= b.X && x <= b.X+b.Width && y >= b.Y && y <= b.Y+b.Height
}
