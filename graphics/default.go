package graphics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ImageResolution = 1024
)

var (
	BrushImage = ebiten.NewImage(1, 1)
)

func init() {
	BrushImage.Fill(color.White)
}
