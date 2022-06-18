package core

import (
	"github.com/Zyko0/Magnet/logic"
	"github.com/Zyko0/Magnet/pkg/geom"
)

const (
	PlayerRadius = 128

	RingRadius          = logic.ScreenHeight / 2
	RingAttractionForce = PlayerFallingVelocity + 0.5
	MaxPlayerDistance   = RingRadius - PlayerRadius

	RingAdvanceSpeed = 0.1
)

type Attraction byte

const (
	AttractionNone Attraction = iota
	AttractionAttract
	AttractionRepel
)

type Ring struct {
	Z          float32
	Center     geom.Vec2
	Attraction Attraction
}

func newRing() *Ring {
	return &Ring{
		Z: 0,
		Center: geom.Vec2{
			X: logic.CenterX,
			Y: logic.CenterY,
		},
		Attraction: AttractionAttract,
	}
}

func (r *Ring) Update() {
	r.Z += 0.01
}

func (r *Ring) GetPlayerRingVelocity(p *Player) geom.Vec2 {
	var sign float32

	switch r.Attraction {
	case AttractionNone:
		return geom.Vec2{}
	case AttractionAttract:
		sign = 1
	case AttractionRepel:
		sign = -1
	}

	force := p.Position.DistanceTo(r.Center) / MaxPlayerDistance

	v := p.Position
	v.Sub(r.Center)
	v.Normalize()
	v.MulN(sign * force * RingAttractionForce)

	return v
}
