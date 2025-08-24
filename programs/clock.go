package programs

import (
	"image"
	"image/color"
	"strings"
	"time"

	"github.com/ducthuy-ng/simple-clock/text"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type ClockProgram struct {
	buffer screen.Buffer
}

func NewClockProgram() *ClockProgram {
	return &ClockProgram{}
}

func (clockProgram *ClockProgram) Init(buffer screen.Buffer) {
	clockProgram.buffer = buffer
}

func (clockProgram *ClockProgram) Draw() screen.Buffer {
	currentTime := time.Now()
	displayTime := time.Now().Format(time.Kitchen)

	if (currentTime.Second() % 2) == 0 {
		displayTime = strings.Replace(displayTime, ":", " ", 1)
	}
	clockProgram.fillBackground(color.RGBA{255, 255, 255, 255})
	clockProgram.renderText(
		displayTime,
		clockProgram.buffer,
		text.DefaultRenderTextOpts().
			SetLocation(fixed.Point26_6{X: fixed.I(clockProgram.buffer.Size().X / 2), Y: fixed.I(clockProgram.buffer.Size().Y / 2)}).
			SetAlign(text.ALIGN_CENTER).
			SetVerticalAlign(text.VERTICAL_ALIGN_CENTER),
	)
	return clockProgram.buffer
}

func (clockProgram *ClockProgram) renderText(input string, buffer screen.Buffer, opts text.RenderTextOpts) {
	// Set text background
	textColorMask := image.NewUniform(opts.Color)

	// Keep it simple, only handle alignment of Left-To-Right text
	fontFace := text.LoadFontFace()
	expectedTextWidth := font.MeasureString(fontFace, input)
	fontHeight := fontFace.Metrics().Height
	drawingDot := opts.Location
	switch opts.Align {
	case text.ALIGN_LEFT:
		drawingDot.X = opts.Location.X
	case text.ALIGN_CENTER:
		drawingDot.X = opts.Location.X - (expectedTextWidth / 2)
	case text.ALIGN_RIGHT:
		drawingDot.X = opts.Location.X - expectedTextWidth
	}

	switch opts.VerticalAlign {
	case text.VERTICAL_ALIGN_TOP:
		drawingDot.Y = opts.Location.Y + fontHeight
	case text.VERTICAL_ALIGN_CENTER:
		drawingDot.Y = opts.Location.Y + (fontHeight / 2)
	}

	fontDrawer := font.Drawer{
		Dst:  buffer.RGBA(),
		Src:  textColorMask,
		Face: fontFace,
		Dot:  drawingDot,
	}
	fontDrawer.DrawString(input)
}

func (clockProgram *ClockProgram) fillBackground(color color.Color) {
	for x := 0; x < clockProgram.buffer.Size().X; x++ {
		for y := 0; y < clockProgram.buffer.Size().Y; y++ {
			clockProgram.buffer.RGBA().Set(x, y, color)
		}
	}
}
