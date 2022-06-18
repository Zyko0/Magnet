package core

import (
	"github.com/Zyko0/Magnet/logic"
	"github.com/Zyko0/Magnet/pkg/geom"
)

const (
	PlayerRadius = 128

	RingRadius        = logic.ScreenHeight / 2
	MaxPlayerDistance = RingRadius - PlayerRadius

	RingAdvanceSpeed = 0.1
)

type Attraction byte

const (
	AttractionNeutral Attraction = iota
	AttractionAttract
	AttractionRepel
)

type Ring struct {
	centerForce geom.Vec2

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
		centerForce: geom.Vec2{
			X: 0,
			Y: 0,
		},
	}
}

func (r *Ring) Update() {
	r.Z += 0.1
}

func (r *Ring) GetPlayerRingVelocity(p *Player) geom.Vec2 {
	// TODO: return r.centerForce * DistancePlayerCenter or something
	return geom.Vec2{
		X: r.Center.X - p.Position.X,
		Y: r.Center.Y - p.Position.Y,
	}
}
