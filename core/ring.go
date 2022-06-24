package core

import (
	"math"
	"math/rand"

	"github.com/Zyko0/Magnet/assets"
	"github.com/Zyko0/Magnet/graphics"
	"github.com/Zyko0/Magnet/logic"
	"github.com/Zyko0/Magnet/pkg/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	PlayerRadius = 60

	RingRadius          = logic.ScreenHeight / 2
	RingAttractionForce = PlayerMoveSpeed * 1.05
	MaxPlayerDistance   = RingRadius - 96 // 96 was originally the player's radius

	SidesCount = graphics.ImageResolution
	DepthToWin = SidesCount

	RotateTextureZInterval = 12
)

type Attraction byte

const (
	AttractionNone Attraction = iota
	AttractionSouth
	AttractionNorth
)

var (
	DataTexture = ebiten.NewImage(graphics.ImageResolution, graphics.ImageResolution)
)

type Coating struct {
	Surfaces []Attraction
}

func newCoating() *Coating {
	surfaces := make([]Attraction, SidesCount)
	for i := 0; i < SidesCount; i++ {
		surfaces[i] = Attraction(rand.Intn(int(AttractionNorth) + 1))
	}

	pixels := make([]byte, SidesCount*SidesCount*4)
	for i := 0; i < SidesCount; i++ {
		switch surfaces[i] {
		case AttractionNone:
			pixels[i*4] = assets.ColorNone.R
			pixels[i*4+1] = assets.ColorNone.G
			pixels[i*4+2] = assets.ColorNone.B
			pixels[i*4+3] = assets.ColorNone.A
		case AttractionSouth:
			pixels[i*4] = assets.ColorSouth.R
			pixels[i*4+1] = assets.ColorSouth.G
			pixels[i*4+2] = assets.ColorSouth.B
			pixels[i*4+3] = assets.ColorSouth.A
		case AttractionNorth:
			pixels[i*4] = assets.ColorNorth.R
			pixels[i*4+1] = assets.ColorNorth.G
			pixels[i*4+2] = assets.ColorNorth.B
			pixels[i*4+3] = assets.ColorNorth.A
		}
	}
	DataTexture.ReplacePixels(pixels)

	return &Coating{
		Surfaces: surfaces,
	}
}

type Ring struct {
	tick             uint64
	lastTextureIndex int

	Z        float32
	Center   geom.Vec2
	Coating  *Coating
	Texture0 *ebiten.Image
	Texture1 *ebiten.Image
	Texture2 *ebiten.Image
}

func newRing() *Ring {
	return &Ring{
		tick:             0,
		lastTextureIndex: 0,

		Z:        0,
		Center:   logic.Center,
		Coating:  newCoating(),
		Texture0: assets.WallTextures[0],
		Texture1: assets.WallTextures[1],
		Texture2: assets.WallTextures[2],
	}
}

func (r *Ring) getPlayerRingVelocity(p *Player) geom.Vec2 {
	var sign float32

	wallMagnet := r.Coating.Surfaces[int(r.Z)]
	switch {
	case wallMagnet == AttractionNone || p.Attraction == AttractionNone:
		return geom.Vec2{}
	case wallMagnet == p.Attraction:
		sign = -1
	case wallMagnet != p.Attraction:
		sign = 1
	}

	force := p.Position.DistanceTo(r.Center) / MaxPlayerDistance
	force = float32(math.Sin(float64(force/2)*math.Pi) + 1)
	force /= 2

	v := p.Position
	v.Sub(r.Center)
	v.Normalize()
	v.MulN(sign * force * RingAttractionForce)

	return v
}

func (r *Ring) Update(ringAdvanceSpeed float32) {
	r.Z += ringAdvanceSpeed
	// Note: if passed 5 depth from texture rotation, can freely unload current one
	if int(r.Z-5)%RotateTextureZInterval == 0 {
		index := (int(r.Z)/RotateTextureZInterval + 2) % 3
		if index != r.lastTextureIndex {
			texture := assets.WallTextures[rand.Intn(len(assets.WallTextures))]
			switch index {
			case 0:
				r.Texture0 = texture
			case 1:
				r.Texture1 = texture
			case 2:
				r.Texture2 = texture
			}

			r.lastTextureIndex = index
		}
	}

	r.tick++
}

func (r *Ring) GetAttraction() Attraction {
	return r.Coating.Surfaces[int(r.Z)]
}
