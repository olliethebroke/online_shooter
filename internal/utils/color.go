package utils

import (
	"image/color"
	"math/rand"
)

// RandomBrightColor generates a random bright color.
func RandomBrightColor() color.RGBA {
	var r, g, b uint8
	for {
		r = uint8(rand.Intn(256))
		g = uint8(rand.Intn(256))
		b = uint8(rand.Intn(256))

		// check if the color is bright enough
		if (int(r) + int(g) + int(b)) > 382 { // 255 * 3 / 2 = 382.5
			break
		}
	}
	return color.RGBA{R: r, G: g, B: b, A: 255}
}
