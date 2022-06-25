package assets

import (
	"bytes"
	_ "embed"
	"image/color"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	brushImg = ebiten.NewImage(1, 1)
	//go:embed images/scifi_brick.png
	scifiBrickSrc   []byte
	scifiBrickImage *ebiten.Image
	//go:embed images/gold_foil.png
	goldFoilSrc   []byte
	goldFoilImage *ebiten.Image
	//go:embed images/black_variative.png
	blackVariativeSrc   []byte
	blackVariativeImage *ebiten.Image
	//go:embed images/concrete.png
	concreteSrc   []byte
	concreteImage *ebiten.Image
	//go:embed images/sandstone.png
	sandstoneSrc   []byte
	sandstoneImage *ebiten.Image
	//go:embed images/marble.png
	marbleSrc   []byte
	marbleImage *ebiten.Image

	WallTextures []*ebiten.Image
	CursorImage  *ebiten.Image
)

func init() {
	brushImg.Fill(color.White)

	img, err := png.Decode(bytes.NewReader(scifiBrickSrc))
	if err != nil {
		log.Fatal(err)
	}
	scifiBrickImage = ebiten.NewImageFromImage(img)

	img, err = png.Decode(bytes.NewReader(goldFoilSrc))
	if err != nil {
		log.Fatal(err)
	}
	goldFoilImage = ebiten.NewImageFromImage(img)

	img, err = png.Decode(bytes.NewReader(blackVariativeSrc))
	if err != nil {
		log.Fatal(err)
	}
	blackVariativeImage = ebiten.NewImageFromImage(img)

	img, err = png.Decode(bytes.NewReader(concreteSrc))
	if err != nil {
		log.Fatal(err)
	}
	concreteImage = ebiten.NewImageFromImage(img)

	img, err = png.Decode(bytes.NewReader(sandstoneSrc))
	if err != nil {
		log.Fatal(err)
	}
	sandstoneImage = ebiten.NewImageFromImage(img)

	img, err = png.Decode(bytes.NewReader(marbleSrc))
	if err != nil {
		log.Fatal(err)
	}
	marbleImage = ebiten.NewImageFromImage(img)

	WallTextures = []*ebiten.Image{
		scifiBrickImage,
		concreteImage,
		sandstoneImage,
		blackVariativeImage,
		goldFoilImage,
		marbleImage,
	}

	CursorImage = ebiten.NewImage(32, 32)
	CursorImage.DrawTriangles([]ebiten.Vertex{
		{
			DstX:   0,
			DstY:   0,
			ColorR: 0.25,
			ColorG: 0.75,
			ColorB: 0.5,
			ColorA: 0.25,
		},
		{
			DstX:   32,
			DstY:   16,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
		{
			DstX:   0,
			DstY:   32,
			ColorR: 0.25,
			ColorG: 0.75,
			ColorB: 0.5,
			ColorA: 0.25,
		},
	}, []uint16{0, 1, 2}, brushImg, nil)
}
