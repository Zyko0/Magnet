package core

import (
	"github.com/Zyko0/Magnet/assets"
	"github.com/Zyko0/Magnet/logic"
	"github.com/Zyko0/Magnet/pkg/geom"
)

type Player struct {
	Angle    float32
	Position geom.Vec2
	Rotation geom.Vec3
	BonesSet assets.BoneSet
}

func newPlayer() *Player {
	return &Player{
		Angle: 0,
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
	baseRotationsAdd = [][]float32{
		assets.BoneSetFalling:  {0, 0, 0},
		assets.BoneSetSliding:  {0.125, 1.1, 2.5},
		assets.BoneSetDashing:  {0, 0, 0},
		assets.BoneSetBouncing: {0, 0, 0},
	}
	baseRotationsMul = [][]float32{
		assets.BoneSetFalling:  {1, 1, 1},
		assets.BoneSetSliding:  {0, 0, 1},
		assets.BoneSetDashing:  {1, 1, 1},
		assets.BoneSetBouncing: {1, 1, 1},
	}
)

func (p *Player) setRotation(ax, ay, az float32) {
	p.Rotation.X = ax*baseRotationsMul[p.BonesSet][0] + baseRotationsAdd[p.BonesSet][0]
	p.Rotation.Y = ay*baseRotationsMul[p.BonesSet][1] + baseRotationsAdd[p.BonesSet][1]
	p.Rotation.Z = az*baseRotationsMul[p.BonesSet][2] + baseRotationsAdd[p.BonesSet][2]
}

func (p *Player) Update() {
}
