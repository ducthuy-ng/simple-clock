package text

import (
	"image/color"

	"golang.org/x/image/math/fixed"
)

type TextAlign int
type VerticalTextAlign int

const (
	ALIGN_LEFT TextAlign = iota
	ALIGN_CENTER
	ALIGN_RIGHT

	VERTICAL_ALIGN_TOP VerticalTextAlign = iota
	VERTICAL_ALIGN_CENTER
	VERTICAL_ALIGN_BOTTOM
)

type RenderTextOpts struct {
	Location      fixed.Point26_6
	Align         TextAlign
	VerticalAlign VerticalTextAlign
	Color         color.RGBA
}

func DefaultRenderTextOpts() RenderTextOpts {
	return RenderTextOpts{
		Location:      fixed.Point26_6{X: 0, Y: 0},
		Align:         ALIGN_LEFT,
		VerticalAlign: VERTICAL_ALIGN_TOP,
		Color:         color.RGBA{0, 0, 0, 255},
	}
}

// Getter methods
func (opts RenderTextOpts) GetLocation() fixed.Point26_6 {
	return opts.Location
}

func (opts RenderTextOpts) GetAlign() TextAlign {
	return opts.Align
}

func (opts RenderTextOpts) GetVerticalAlign() VerticalTextAlign {
	return opts.VerticalAlign
}

func (opts RenderTextOpts) GetColor() color.RGBA {
	return opts.Color
}

// Setter methods following builder pattern (return modified copy)
func (opts RenderTextOpts) SetLocation(location fixed.Point26_6) RenderTextOpts {
	opts.Location = location
	return opts
}

func (opts RenderTextOpts) SetAlign(align TextAlign) RenderTextOpts {
	opts.Align = align
	return opts
}

func (opts RenderTextOpts) SetVerticalAlign(verticalAlign VerticalTextAlign) RenderTextOpts {
	opts.VerticalAlign = verticalAlign
	return opts
}

func (opts RenderTextOpts) SetColor(c color.RGBA) RenderTextOpts {
	opts.Color = c
	return opts
}
