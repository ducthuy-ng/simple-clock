package main

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed assets
var assetFs embed.FS

const (
	normalFontSize = 60
	fontFilePath   = "assets/Roboto-VariableFont_wdth,wght.ttf"
)

type Clock struct {
	textFaceSource *text.GoTextFaceSource
	currentTime    time.Time
}

func NewClock() *Clock {
	fontFile, err := fs.ReadFile(assetFs, fontFilePath)
	if err != nil {
		panic(fmt.Sprintf("failed to load font: %v", err))
	}
	textFaceSource, err := text.NewGoTextFaceSource(bytes.NewReader(fontFile))
	if err != nil {
		panic(fmt.Sprintf("failed to create clock image: %v", err))
	}
	return &Clock{
		textFaceSource: textFaceSource,
	}
}

func (clock *Clock) Update() error {
	clock.currentTime = time.Now()
	return nil
}

func (clock *Clock) Draw(screen *ebiten.Image) {
	layoutOpts := &text.LayoutOptions{PrimaryAlign: text.AlignCenter}
	drawOpts := &text.DrawOptions{LayoutOptions: *layoutOpts}

	textFaceOpts := &text.GoTextFace{
		Source: clock.textFaceSource,
		Size:   normalFontSize,
	}
	displayTime := clock.currentTime.Format(time.Kitchen)

	drawOpts.GeoM.Translate(float64(screen.Bounds().Dx())/2, float64(screen.Bounds().Dy())/2)

	text.Draw(screen, displayTime, textFaceOpts, drawOpts)
}

func (clock *Clock) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
