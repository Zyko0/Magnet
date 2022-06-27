package ui

import (
	"image/color"

	"github.com/Zyko0/Magnet/graphics"
	"github.com/Zyko0/Magnet/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameOverView struct {
	bgImage *ebiten.Image
}

func NewGameOver() *GameOverView {
	const (
		borderWidth       = 4
		cardWidth         = logic.ScreenHeight / 2
		noBorderCardWidth = logic.ScreenHeight/2 - borderWidth*2
	)

	bgImage := ebiten.NewImage(logic.ScreenHeight/2, logic.ScreenHeight/2)
	bgImage.Fill(color.White)

	vertices, indices := graphics.AppendQuadVerticesIndices(
		nil, nil,
		borderWidth, borderWidth,
		noBorderCardWidth, noBorderCardWidth,
		0, 0, 0, 1, 0,
	)
	bgImage.DrawTriangles(vertices, indices, graphics.BrushImage, nil)

	return &GameOverView{
		bgImage: bgImage,
	}
}

func (g *GameOverView) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(logic.ScreenWidth/2-g.bgImage.Bounds().Dx()/2),
		float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2),
	)
	screen.DrawImage(g.bgImage, op)

	/*
		// Title
		rect := text.BoundString(assets.DefaultFontFace, g.title)
		geom := ebiten.GeoM{}
		geom.Translate(
			float64(logic.ScreenWidth/2-rect.Max.X/2),
			float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36,
		)
		text.DrawWithOptions(screen, g.title, assets.DefaultFontFace, &ebiten.DrawImageOptions{
			GeoM: geom,
		})
		// Completion text
		rect = text.BoundString(assets.DefaultSmallFontFace, g.completion)
		geom = ebiten.GeoM{}
		geom.Translate(
			float64(logic.ScreenWidth/2-rect.Max.X/2),
			float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36*2,
		)
		text.DrawWithOptions(screen, g.completion, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
			GeoM: geom,
		})
		// Duration text
		rect = text.BoundString(assets.DefaultSmallFontFace, g.duration)
		geom = ebiten.GeoM{}
		geom.Translate(
			float64(logic.ScreenWidth/2-rect.Max.X/2),
			float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36*3,
		)
		text.DrawWithOptions(screen, g.duration, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
			GeoM: geom,
		})
		// Loop count text
		rect = text.BoundString(assets.DefaultSmallFontFace, g.loopCount)
		geom = ebiten.GeoM{}
		geom.Translate(
			float64(logic.ScreenWidth/2-rect.Max.X/2),
			float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36*4,
		)
		text.DrawWithOptions(screen, g.loopCount, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
			GeoM: geom,
		})
		// HP text
		rect = text.BoundString(assets.DefaultSmallFontFace, g.hp)
		geom = ebiten.GeoM{}
		geom.Translate(
			float64(logic.ScreenWidth/2-rect.Max.X/2),
			float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36*5,
		)
		text.DrawWithOptions(screen, g.hp, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
			GeoM: geom,
		})
		// Press R to restart
		const str = "Press <Backspace> to restart"

		rect = text.BoundString(assets.DefaultSmallFontFace, str)
		geom = ebiten.GeoM{}
		geom.Translate(
			float64(logic.ScreenWidth/2-rect.Max.X/2),
			float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36*6.5,
		)
		text.DrawWithOptions(screen, str, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
			GeoM: geom,
		})
	*/
}
