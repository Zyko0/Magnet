package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/Zyko0/Magnet/assets"
	"github.com/Zyko0/Magnet/core"
	"github.com/Zyko0/Magnet/graphics"
	"github.com/Zyko0/Magnet/logic"
	"github.com/Zyko0/Magnet/pkg/geom"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Game struct {
	tick   int
	paused bool

	game *core.Game
}

func New() *Game {
	return &Game{
		game: core.NewGame(),
	}
}

func (g *Game) Update() error {
	g.tick++

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("quit")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.paused = !g.paused
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
		for _, t := range g.game.Obstacles[i].Triangles {
			vertices, indices = t.AppendVerticesIndices(vertices, indices, index,
				1, 1, 1, 1,
			)
			index++
		}
	}
	for i := range vertices {
		vertices[i].SrcX = geom.Clamp(vertices[i].DstX-assets.ShapeOffsetX, 0, assets.ShapeSize)
		vertices[i].SrcY = geom.Clamp(vertices[i].DstY, 0, assets.ShapeSize)
	}
	screen.DrawTriangles(vertices, indices, assets.ShapeCircleMaskImage, nil)

	// Ring
	vertices, indices = graphics.AppendQuadVerticesIndices(vertices[:0], indices[:0],
		logic.ScreenWidth/2-logic.ScreenHeight/2, 0, logic.ScreenHeight, logic.ScreenHeight,
		1, 1, 1, 1, 0,
	)
	screen.DrawTrianglesShader(vertices, indices, assets.RingShader, &ebiten.DrawTrianglesShaderOptions{
		Uniforms: map[string]interface{}{
			"Time": float32(g.tick) / logic.TPS,
		},
	})

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

	// TODO: Player bounding circle debug
	vertices, indices = graphics.AppendQuadVerticesIndices(vertices[:0], indices,
		x-core.PlayerRadius, y-core.PlayerRadius, core.PlayerRadius*2, core.PlayerRadius*2,
		1, 1, 1, 1, 0,
	)
	screen.DrawTrianglesShader(vertices, indices, assets.RingShader, &ebiten.DrawTrianglesShaderOptions{
		Uniforms: map[string]interface{}{
			"Time": float32(g.tick) / logic.TPS,
		},
	})

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS %.2f - FPS %.2f - Z: %.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS(), g.game.Ring.Z))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return logic.ScreenWidth, logic.ScreenHeight
}

func main() {
	// Note: force opengl
	os.Setenv("EBITEN_GRAPHICS_LIBRARY", "opengl")

	ebiten.SetFullscreen(true)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum) // TODO: vsync on
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	ebiten.SetWindowResizable(true)
	ebiten.SetMaxTPS(logic.TPS)
	if err := ebiten.RunGame(New()); err != nil {
		log.Fatal(err)
	}
}
