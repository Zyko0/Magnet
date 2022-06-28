package assets

import (
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomonobold"
)

var (
	DefaultFontFace      font.Face
	DefaultSmallFontFace font.Face
)

func init() {
	pfont, err := truetype.Parse(gomonobold.TTF)
	if err != nil {
		log.Fatal(err)
	}

	DefaultFontFace = truetype.NewFace(pfont, &truetype.Options{
		Size: 24,
	})
	DefaultSmallFontFace = truetype.NewFace(pfont, &truetype.Options{
		Size: 18,
	})
}
