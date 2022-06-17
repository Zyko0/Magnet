package assets

import (
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed shaders/player.kage
	playerShaderSrc []byte
	PlayerShader    *ebiten.Shader

	//go:embed shaders/tunnel.kage
	tunnelShaderSrc []byte
	TunnelShader    *ebiten.Shader

	//go:embed shaders/ring.kage
	ringShaderSrc []byte
	RingShader    *ebiten.Shader
)

func init() {
	var err error

	PlayerShader, err = ebiten.NewShader(playerShaderSrc)
	if err != nil {
		log.Fatal(err)
	}

	TunnelShader, err = ebiten.NewShader(tunnelShaderSrc)
	if err != nil {
		log.Fatal(err)
	}

	RingShader, err = ebiten.NewShader(ringShaderSrc)
	if err != nil {
		log.Fatal(err)
	}
}
