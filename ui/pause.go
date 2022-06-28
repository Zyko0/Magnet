package ui

import (
	"image/color"

	"github.com/Zyko0/Magnet/assets"
	"github.com/Zyko0/Magnet/graphics"
	"github.com/Zyko0/Magnet/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type PauseView struct {
	bgImage *ebiten.Image
}

func NewPauseView() *PauseView {
	const (
		borderWidth        = 4
		cardWidth          = logic.ScreenHeight / 2
		cardHeight         = logic.ScreenHeight / 6
		noBorderCardWidth  = cardWidth - borderWidth*2
		noBorderCardHeight = cardHeight - borderWidth*2
	)

	bgImage := ebiten.NewImage(cardWidth, cardHeight)
	bgImage.Fill(color.White)

	vertices, indices := graphics.AppendQuadVerticesIndices(
		nil, nil,
		borderWidth, borderWidth,
		noBorderCardWidth, noBorderCardHeight,
		0, 0, 0, 1, 0,
	)
	bgImage.DrawTriangles(vertices, indices, graphics.BrushImage, nil)

	// Title
	const gameOverStr = "Pause"
	rect := text.BoundString(assets.DefaultFontFace, gameOverStr)
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(bgImage.Bounds().Dx()/2-rect.Max.X/2),
		float64(rect.Dy())+36,
	)
	text.DrawWithOptions(bgImage, gameOverStr, assets.DefaultFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// Instruction text
	const restartStr = "Press <P> to resume, <R> to restart"
	rect = text.BoundString(assets.DefaultSmallFontFace, restartStr)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(bgImage.Bounds().Dx()/2-rect.Max.X/2),
		float64(rect.Dy())+36*3,
	)
	text.DrawWithOptions(bgImage, restartStr, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})

	return &PauseView{
		bgImage: bgImage,
	}
}

func (p *PauseView) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(logic.ScreenWidth/2-p.bgImage.Bounds().Dx()/2),
		float64(logic.ScreenHeight/2-p.bgImage.Bounds().Dy()/2),
	)
	op.ColorM.Scale(1, 1, 1, 0.75)
	screen.DrawImage(p.bgImage, op)
}
