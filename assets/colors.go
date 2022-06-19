package assets

import "image/color"

var (
	ColorNone  = color.RGBA{0, 0, 0, 0}
	ColorSouth = color.RGBA{
		R: 255,
		G: 16,
		B: 0,
		A: 255,
	}
	ColorNorth = color.RGBA{
		R: 51,
		G: 0,
		B: 128,
		A: 255,
	}
)
