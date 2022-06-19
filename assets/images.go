package assets

import (
	"bytes"
	_ "embed"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed images/scifi_brick.png
	scifiBrickSrc   []byte
	scifiBrickImage *ebiten.Image
	//go:embed images/aluminium_hexagon.png
	aluminiumHexagonSrc   []byte
	aluminiumHexagonImage *ebiten.Image
	//go:embed images/moroccan_hexagon.png
	moroccanHexagonSrc   []byte
	moroccanHexagonImage *ebiten.Image
	//go:embed images/retro_futuristic.png
	retroFuturisticSrc   []byte
	retroFuturisticImage *ebiten.Image
	//go:embed images/gold_foil.png
	goldFoilSrc   []byte
	goldFoilImage *ebiten.Image

	WallTextures []*ebiten.Image
)

func init() {
	img, err := png.Decode(bytes.NewReader(scifiBrickSrc))
	if err != nil {
		log.Fatal(err)
	}
	scifiBrickImage = ebiten.NewImageFromImage(img)

	img, err = png.Decode(bytes.NewReader(aluminiumHexagonSrc))
	if err != nil {
		log.Fatal(err)
	}
	aluminiumHexagonImage = ebiten.NewImageFromImage(img)

	img, err = png.Decode(bytes.NewReader(moroccanHexagonSrc))
	if err != nil {
		log.Fatal(err)
	}
	moroccanHexagonImage = ebiten.NewImageFromImage(img)

	img, err = png.Decode(bytes.NewReader(retroFuturisticSrc))
	if err != nil {
		log.Fatal(err)
	}
	retroFuturisticImage = ebiten.NewImageFromImage(img)

	img, err = png.Decode(bytes.NewReader(goldFoilSrc))
	if err != nil {
		log.Fatal(err)
	}
	goldFoilImage = ebiten.NewImageFromImage(img)

	WallTextures = []*ebiten.Image{
		scifiBrickImage,
		//aluminiumHexagonImage,
		//moroccanHexagonImage,
		retroFuturisticImage,
		goldFoilImage,
	}
}
