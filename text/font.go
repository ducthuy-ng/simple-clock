package text

import (
	"log"

	"github.com/ducthuy-ng/simple-clock/assets"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func LoadFontFace() font.Face {
	fontFile, err := assets.Assets.ReadFile("Roboto-VariableFont_wdth,wght.ttf")
	if err != nil {
		log.Panicf("failed to load font: %v", err)
	}

	font, err := opentype.Parse(fontFile)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}
	fontFace, err := opentype.NewFace(font, &opentype.FaceOptions{Size: 48, DPI: 72})
	if err != nil {
		log.Fatalf("failed to create font face: %v", err)
	}
	return fontFace
}
