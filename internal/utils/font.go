package utils

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"io"
	"os"
)

// LoadFont loads a font from the file.
//
// Accepts a string path to the file.
//
// Returns an implementation of the font.Face
// interface and an error if it exists
// otherwise returns nil.
func LoadFont(path string) (font.Face, error) {
	// open the font file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read the font data
	fontData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// parse the font
	ttf, err := opentype.Parse(fontData)
	if err != nil {
		return nil, err
	}

	// create a font face object
	fontFace, err := opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}

	return fontFace, nil
}
