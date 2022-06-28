package main

import (
	"errors"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Zyko0/Magnet/assets"
	"github.com/Zyko0/Magnet/core"
	"github.com/Zyko0/Magnet/graphics"
	"github.com/Zyko0/Magnet/logic"
	"github.com/Zyko0/Magnet/pkg/geom"
	"github.com/Zyko0/Magnet/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Game struct {
	tick   int
	paused bool

	splashView   *ui.SplashView
	gameOverView *ui.GameOverView
	pauseView    *ui.PauseView
	game         *core.Game
}

func New() *Game {
	return &Game{
		gameOverView: ui.NewGameOver(),
		splashView:   ui.NewSplashView(),
		pauseView:    ui.NewPauseView(),
		game:         core.NewGame(),
	}
}

func (g *Game) Update() error {
	g.tick++

	// Refresh inputs if on mobile
	logic.UpdateInputs()
	// TODO: remove
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("quit")
	}

	if g.splashView.Active() {
		g.splashView.Update()
		// If still active, just abort
		if g.splashView.Active() {
			return nil
		}
		// If it has been deactivated just there, start a game
		g.game = core.NewGame()
		assets.ReplayGameMusic()
	}

	// Restart
	if inpututil.IsKeyJustPressed(ebiten.KeyR) || (g.game.Over && len(logic.Touches) > 0) {
		g.game = core.NewGame()
		g.paused = false
		assets.StopSlideSound()
		assets.ReplayGameMusic()
	}

	if g.game.Over {
		assets.StopSlideSound()
		assets.StopGameMusic()
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.paused = !g.paused
		if g.paused {
			assets.StopSlideSound()
			assets.StopGameMusic()
		} else {
			assets.ResumeGameMusic()
		}
	}
	if g.paused {
		return nil
	}

	g.game.Update()

	return nil
}

var (
	// Note: :)
	magicDepthCorrection = float32(math.Sqrt(1. / 3.))
)

func (g *Game) Draw(screen *ebiten.Image) {
	const (
		PlayerSize = 512
	)

	// Only draw splash view if still active
	if g.splashView.Active() {
		g.splashView.Draw(screen)
		return
	}

	// Tunnel
	vertices, indices := graphics.AppendQuadVerticesIndices(nil, nil,
		0, 0, logic.ScreenWidth, logic.ScreenHeight,
		1, 1, 1, 1, 0,
	)
	for i := range vertices {
		vertices[i].SrcX *= graphics.ImageResolution
		vertices[i].SrcY *= graphics.ImageResolution
	}
	screen.DrawTrianglesShader(vertices, indices, assets.TunnelShader, &ebiten.DrawTrianglesShaderOptions{
		Uniforms: map[string]interface{}{
			"Depth":                  g.game.Ring.Z - magicDepthCorrection,
			"RotateTextureZInterval": float32(core.RotateTextureZInterval),
			"LightPosition": []float32{
				g.game.Player.Position.X,
				g.game.Player.Position.Y,
			},
		},
		Images: [4]*ebiten.Image{
			core.DataTexture,
			g.game.Ring.Texture0,
			g.game.Ring.Texture1,
			g.game.Ring.Texture2,
		},
	})
	// Obstacles
	vertices, indices = vertices[:0], indices[:0]
	index := 0
	for i := len(g.game.Obstacles) - 1; i >= 0; i-- {
		o := g.game.Obstacles[i]
		clr := o.GetColor()
		alpha := o.GetAlpha()
		for _, t := range o.Triangles {
			vertices, indices = t.AppendVerticesIndices(vertices, indices, index,
				clr[0]*o.Scale,
				clr[1]*o.Scale,
				clr[2]*o.Scale,
				alpha,
			)
			index++
		}
	}
	// Update source image coordinates
	index = 0
	for i := len(g.game.Obstacles) - 1; i >= 0; i-- {
		for _, t := range g.game.Obstacles[i].SrcTriangles {
			vertices[index*3+0].SrcX = geom.Clamp(t.A.X-assets.ShapeOffsetX, 0, assets.ShapeSize)
			vertices[index*3+0].SrcY = geom.Clamp(t.A.Y, 0, assets.ShapeSize)
			vertices[index*3+1].SrcX = geom.Clamp(t.B.X-assets.ShapeOffsetX, 0, assets.ShapeSize)
			vertices[index*3+1].SrcY = geom.Clamp(t.B.Y, 0, assets.ShapeSize)
			vertices[index*3+2].SrcX = geom.Clamp(t.C.X-assets.ShapeOffsetX, 0, assets.ShapeSize)
			vertices[index*3+2].SrcY = geom.Clamp(t.C.Y, 0, assets.ShapeSize)
			index++
		}
	}
	screen.DrawTriangles(vertices, indices, assets.ShapeCircleMaskImage, nil)

	// Ring
	vertices, indices = graphics.AppendQuadVerticesIndices(vertices[:0], indices[:0],
		logic.ScreenWidth/2-logic.ScreenHeight/2, 0, logic.ScreenHeight, logic.ScreenHeight,
		0.2, 0.2, 0.2, 0.2, 0,
	)
	screen.DrawTrianglesShader(vertices, indices, assets.RingShader, nil)

	// Player
	x, y := g.game.Player.Position.X, g.game.Player.Position.Y
	vertices, indices = graphics.AppendQuadVerticesIndices(vertices[:0], indices[:0],
		float32(x)-PlayerSize/2, float32(y)-PlayerSize/2, PlayerSize, PlayerSize,
		1, 1, 1, 1, 0,
	)

	r := g.game.Player.Rotation
	uniforms := map[string]interface{}{
		"Rotation": []float32{r.X, r.Y, r.Z},
		"Color":    g.game.Player.GetColor(),
	}
	uniforms = g.game.Player.BonesSet.GetBones().AppendUniforms(uniforms)
	screen.DrawTrianglesShader(vertices, indices, assets.PlayerShader, &ebiten.DrawTrianglesShaderOptions{
		Uniforms: uniforms,
	})
	// Player dash energy bar if necessary
	if g.game.Player.DashEnergy > 0 && g.game.Player.DashEnergy < 1 {
		x0, y0 := g.game.Player.Position.X-core.PlayerRadius*1.5, g.game.Player.Position.Y+core.PlayerRadius*2
		w, h := float32(core.PlayerRadius*3), float32(24)
		// Bar
		graphics.DrawRect(screen,
			x0, y0, w*g.game.Player.DashEnergy, h,
			0.8, 0, 0.8, 0.5,
		)
		// Border
		graphics.DrawRectBorder(screen,
			x0, y0, w, h, 1,
			1, 1, 1, 0.7,
		)
	}

	// If it's game over we have nothing more to draw
	if g.game.Over {
		g.gameOverView.Draw(screen, g.game.GetScore(), g.game.GetDifficulty().String())
		return
	}
	// Display score
	scoreStr := strconv.FormatInt(g.game.GetScore(), 10)
	rect := text.BoundString(assets.DefaultFontFace, scoreStr)
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(logic.ScreenWidth/2-rect.Max.X/2), 128,
	)
	colorM := ebiten.ColorM{}
	colorM.Scale(1, 1, 1, 0.5)
	text.DrawWithOptions(screen, scoreStr, assets.DefaultFontFace, &ebiten.DrawImageOptions{
		GeoM:   geom,
		ColorM: colorM,
	})
	// If paused we just draw pause card
	if g.paused {
		g.pauseView.Draw(screen)
	}

	// Draw cursor
	if !g.game.Direction.IsZero() {
		cx, cy := ebiten.CursorPosition()
		geom := ebiten.GeoM{}
		geom.Rotate(float64(g.game.Direction.Atan2()))
		geom.Translate(float64(cx), float64(cy))
		screen.DrawImage(assets.CursorImage, &ebiten.DrawImageOptions{
			GeoM: geom,
		})
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return logic.ScreenWidth, logic.ScreenHeight
}

func main() {
	// Note: force opengl since directx is not stable yet
	os.Setenv("EBITEN_GRAPHICS_LIBRARY", "opengl")

	ebiten.SetFullscreen(true)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	ebiten.SetMaxTPS(logic.TPS)
	if err := ebiten.RunGame(New()); err != nil {
		log.Fatal(err)
	}
}
