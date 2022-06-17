package assets

import (
	"bytes"
	_ "embed"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed textures/scifi_brick.png
	scifiBrickSrc   []byte
	ScifiBrickImage *ebiten.Image
	//go:embed textures/aluminium_hexagon.png
	aluminiumHexagonSrc   []byte
	AluminiumHexagonImage *ebiten.Image
	//go:embed textures/moroccan_hexagon.png
	moroccanHexagonSrc   []byte
	MoroccanHexagonImage *ebiten.Image
	//go:embed textures/retro_futuristic.png
	retroFuturisticSrc   []byte
	RetroFuturisticImage *ebiten.Image
	//go:embed textures/gold_foil.png
	goldFoilSrc   []byte
	GoldFoilImage *ebiten.Image
)

func init() {
	img, err := png.Decode(bytes.NewReader(scifiBrickSrc))
	if err != nil {
		log.Fatal(err)
	}
	ScifiBrickImage = ebiten.NewImageFromImage(img)

	img, err = png.Decode(bytes.NewReader(aluminiumHexagonSrc))
	if err != nil {
		log.Fatal(err)
	}
	AluminiumHexagonImage = ebiten.NewImageFromImage(img)

	img, err = png.Decode(bytes.NewReader(moroccanHexagonSrc))
	if err != nil {
		log.Fatal(err)
	}
	MoroccanHexagonImage = ebiten.NewImageFromImage(img)

	img, err = png.Decode(bytes.NewReader(retroFuturisticSrc))
	if err != nil {
		log.Fatal(err)
	}
	RetroFuturisticImage = ebiten.NewImageFromImage(img)

	img, err = png.Decode(bytes.NewReader(goldFoilSrc))
	if err != nil {
		log.Fatal(err)
	}
	GoldFoilImage = ebiten.NewImageFromImage(img)
}
