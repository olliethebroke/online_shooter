package assets

import (
	"golang.org/x/image/font"
	"online_shooter/internal/logger"
	"online_shooter/internal/utils"
)

const pathToAssets = "./internal/assets/"
const fontFile = "fonts/minecraft.ttf"

type Assets struct {
	font font.Face
}

var assets = Assets{}

// Font returns a font from the assets.
// If the font is not initialized
// methods loads it from the assets folder.
//
// Returns the font.Face interface implementation.
func Font() font.Face {
	if assets.font == nil {
		f, err := utils.LoadFont(pathToAssets + fontFile)
		if err != nil {
			logger.Warn(err)
		}

		assets.font = f
	}

	return assets.font
}
