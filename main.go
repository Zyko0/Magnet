package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/Zyko0/Magnet/assets"
	"github.com/Zyko0/Magnet/core"
	"github.com/Zyko0/Magnet/graphics"
	"github.com/Zyko0/Magnet/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Game struct {
	tick int

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

	g.game.Update()

	return nil
}

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
		vertices[i].SrcX *= float32(assets.MoroccanHexagonImage.Bounds().Dx())
		vertices[i].SrcY *= float32(assets.MoroccanHexagonImage.Bounds().Dy())
	}
	screen.DrawTrianglesShader(vertices, indices, assets.TunnelShader, &ebiten.DrawTrianglesShaderOptions{
		Uniforms: map[string]interface{}{
			"Time": float32(g.tick) / logic.TPS,
		},
		Images: [4]*ebiten.Image{
			assets.ScifiBrickImage,
		},
	})

	// Ring
	vertices, indices = graphics.AppendQuadVerticesIndices(nil, nil,
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
	vertices, indices = graphics.AppendQuadVerticesIndices(nil, nil,
		float32(x)-PlayerSize/2, float32(y)-PlayerSize/2, PlayerSize, PlayerSize,
		1, 1, 1, 1, 0,
	)

	r := g.game.Player.Rotation
	uniforms := map[string]interface{}{
		"Rotation": []float32{r.X, r.Y, r.Z},
	}
	uniforms = g.game.Player.BonesSet.GetBones().AppendUniforms(uniforms)
	screen.DrawTrianglesShader(vertices, indices, assets.PlayerShader, &ebiten.DrawTrianglesShaderOptions{
		Uniforms: uniforms,
	})

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS %.2f - FPS %.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
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
