package core

import (
	"github.com/Zyko0/Magnet/assets"
	"github.com/Zyko0/Magnet/logic"
	"github.com/Zyko0/Magnet/pkg/geom"
)

type Player struct {
	DashingTicks uint64
	DashEnergy   float32
	Attraction   Attraction
	Position     geom.Vec2
	Rotation     geom.Vec3
	BonesSet     assets.BoneSet
}

func newPlayer() *Player {
	return &Player{
		DashingTicks: 0,
		DashEnergy:   1,
		Attraction:   AttractionNone,
		Position: geom.Vec2{
			X: logic.ScreenWidth / 2,
			Y: logic.ScreenHeight / 2,
		},
		Rotation: geom.Vec3{
			X: 0,
			Y: 0,
			Z: 0,
		},
		BonesSet: assets.BoneSetSliding,
	}
}

var (
	baseRotationsAdd = []geom.Vec3{
		assets.BoneSetFalling: {X: 0, Y: 0, Z: 0},
		assets.BoneSetSliding: {X: 0.125, Y: 1.1, Z: 2.5},
		assets.BoneSetDashing: {X: 0, Y: 2.5, Z: -1.3},
	}
	baseRotationsMul = []geom.Vec3{
		assets.BoneSetFalling: {X: 1, Y: 1, Z: 0},
		assets.BoneSetSliding: {X: 0, Y: 0, Z: 1},
		assets.BoneSetDashing: {X: -0.25, Y: -0.25, Z: 1},
	}
)

func (p *Player) setRotation(ax, ay, az float32) {
	p.Rotation.X = ax*baseRotationsMul[p.BonesSet].X + baseRotationsAdd[p.BonesSet].X
	p.Rotation.Y = ay*baseRotationsMul[p.BonesSet].Y + baseRotationsAdd[p.BonesSet].Y
	p.Rotation.Z = az*baseRotationsMul[p.BonesSet].Z + baseRotationsAdd[p.BonesSet].Z
}

var (
	colorsByAttraction = [][]float32{
		AttractionNone: {
			0.5, 0.5, 0.5,
		},
		AttractionSouth: {
			float32(assets.ColorSouth.R) / 255,
			float32(assets.ColorSouth.G) / 255,
			float32(assets.ColorSouth.B) / 255,
		},
		AttractionNorth: {
			float32(assets.ColorNorth.R) / 255,
			float32(assets.ColorNorth.G) / 255,
			float32(assets.ColorNorth.B) / 255,
		},
	}
)

func (p *Player) GetDashMultiplier() float32 {
	const ticksToReachMaxDash = logic.TPS / 2

	ticks := p.DashingTicks
	if p.DashingTicks > logic.TPS/2 {
		ticks = logic.TPS / 2
	}

	return 1 + float32(ticks)/ticksToReachMaxDash
}

func (p *Player) GetColor() []float32 {
	return colorsByAttraction[p.Attraction]
}

func (p *Player) Update() {
	const (
		dashEnergyFillRate        = 0.005
		dashEnergyConsumptionRate = dashEnergyFillRate * 3
	)

	if logic.IntentDash {
		if p.DashEnergy > 0 {
			p.DashingTicks++
			p.DashEnergy = geom.Clamp(p.DashEnergy-dashEnergyConsumptionRate, 0, 1)
			p.BonesSet = assets.BoneSetDashing
			if p.DashingTicks == 1 {
				assets.PlayDashSound()
			}
		} else {
			p.DashingTicks = 0
		}
	} else {
		p.DashingTicks = 0
		p.BonesSet = assets.BoneSetFalling
		p.DashEnergy = geom.Clamp(p.DashEnergy+dashEnergyFillRate, 0, 1)
		assets.StopDashSound()
	}
}
