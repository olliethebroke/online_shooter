package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"math/rand"
)

// downloadImage creates ebiten image from file.
//
// Accepts string file name of an image as an argument.
//
// Returns a pointer to ebiten.Image.
func downloadImage(imgName string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(imgName)
	if err != nil {
		log.Fatal(img, " - failed to create image from file: ", err)
	}
	return img
}

// randomBrightColor generates a random bright color.
func randomBrightColor() color.RGBA {
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
