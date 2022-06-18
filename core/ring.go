package core

import (
	"math"
	"math/rand"

	"github.com/Zyko0/Magnet/graphics"
	"github.com/Zyko0/Magnet/logic"
	"github.com/Zyko0/Magnet/pkg/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	PlayerRadius = 128

	RingRadius          = logic.ScreenHeight / 2
	RingAttractionForce = PlayerFallingVelocity
	MaxPlayerDistance   = RingRadius - PlayerRadius

	RingAdvanceSpeed = 0.01

	SidesCount = graphics.ImageResolution
)

type Attraction byte

const (
	AttractionNone Attraction = iota
	AttractionAttract
	AttractionRepel
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
		surfaces[i] = Attraction(rand.Intn(int(AttractionRepel) + 1))
	}

	pixels := make([]byte, SidesCount*SidesCount*4)
	for i := 0; i < SidesCount; i++ {
		switch surfaces[i] {
		case AttractionNone:
			pixels[i*4] = 0
			pixels[i*4+1] = 0
			pixels[i*4+2] = 0
			pixels[i*4+3] = 255
		case AttractionAttract:
			pixels[i*4] = 0
			pixels[i*4+1] = 255
			pixels[i*4+2] = 0
			pixels[i*4+3] = 255
		case AttractionRepel:
			pixels[i*4] = 255
			pixels[i*4+1] = 0
			pixels[i*4+2] = 0
			pixels[i*4+3] = 255
		}
	}
	DataTexture.ReplacePixels(pixels)

	return &Coating{
		Surfaces: surfaces,
	}
}

type Ring struct {
	Z       float32
	Center  geom.Vec2
	Coating *Coating
}

func newRing() *Ring {
	return &Ring{
		Z: 0,
		Center: geom.Vec2{
			X: logic.CenterX,
			Y: logic.CenterY,
		},
		Coating: newCoating(),
	}
}

func (r *Ring) getPlayerRingVelocity(p *Player) geom.Vec2 {
	var sign float32

	switch r.Coating.Surfaces[int(r.Z)] {
	case AttractionNone:
		return geom.Vec2{}
	case AttractionAttract:
		sign = 1
	case AttractionRepel:
		sign = -1
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

func (r *Ring) Update() {
	r.Z += RingAdvanceSpeed
}

func (r *Ring) GetAttraction() Attraction {
	return r.Coating.Surfaces[int(r.Z)]
}
