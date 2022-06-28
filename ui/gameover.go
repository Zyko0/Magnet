package ui

import (
	"image/color"
	"strconv"

	"github.com/Zyko0/Magnet/assets"
	"github.com/Zyko0/Magnet/core"
	"github.com/Zyko0/Magnet/graphics"
	"github.com/Zyko0/Magnet/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type GameOverView struct {
	bgImage *ebiten.Image
}

func NewGameOver() *GameOverView {
	const (
		borderWidth        = 4
		cardWidth          = logic.ScreenHeight / 3 * 2
		cardHeight         = logic.ScreenHeight / 4
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

	return &GameOverView{
		bgImage: bgImage,
	}
}

func (g *GameOverView) Draw(screen *ebiten.Image, score int64, difficulty string) {
	colorM := ebiten.ColorM{}
	colorM.Scale(1, 1, 1, 0.75)

	op := &ebiten.DrawImageOptions{
		ColorM: colorM,
	}
	op.GeoM.Translate(
		float64(logic.ScreenWidth/2-g.bgImage.Bounds().Dx()/2),
		float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2),
	)
	screen.DrawImage(g.bgImage, op)

	// Title
	const gameOverStr = "Game Over"
	rect := text.BoundString(assets.DefaultFontFace, gameOverStr)
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(logic.ScreenWidth/2-rect.Max.X/2),
		float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36,
	)
	text.DrawWithOptions(screen, gameOverStr, assets.DefaultFontFace, &ebiten.DrawImageOptions{
		GeoM:   geom,
		ColorM: colorM,
	})
	// Score text
	scoreStr := "Score: " + strconv.FormatInt(score, 10)
	rect = text.BoundString(assets.DefaultSmallFontFace, scoreStr)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(logic.ScreenWidth/2-rect.Max.X/2),
		float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36*2,
	)
	text.DrawWithOptions(screen, scoreStr, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
		GeoM:   geom,
		ColorM: colorM,
	})
	// Difficulty text
	difficultyReachedStr := "You reached " + difficulty + " difficulty"
	rect = text.BoundString(assets.DefaultSmallFontFace, difficultyReachedStr)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(logic.ScreenWidth/2-rect.Max.X/2),
		float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36*3,
	)
	text.DrawWithOptions(screen, difficultyReachedStr, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
		GeoM:   geom,
		ColorM: colorM,
	})
	// Instruction text
	const restartStr = "Press <R> / Tap the screen to restart"
	rect = text.BoundString(assets.DefaultSmallFontFace, restartStr)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(logic.ScreenWidth/2-rect.Max.X/2),
		float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36*5,
	)
	text.DrawWithOptions(screen, restartStr, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
		GeoM:   geom,
		ColorM: colorM,
	})
	// Tip text
	tip := core.GetTip()
	rect = text.BoundString(assets.DefaultSmallFontFace, tip)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(logic.ScreenWidth/2-rect.Max.X/2),
		float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36*6,
	)
	text.DrawWithOptions(screen, tip, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
		GeoM:   geom,
		ColorM: colorM,
	})
}
