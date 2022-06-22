package logic

import "github.com/Zyko0/Magnet/pkg/geom"

const (
	TPS                       = 60
	ScreenWidth, ScreenHeight = 1920, 1080
	CenterX, CenterY          = ScreenWidth / 2, ScreenHeight / 2
)

var (
	Center = geom.Vec2{
		X: CenterX,
		Y: CenterY,
	}
)
