package graphics

import "github.com/Zyko0/Magnet/pkg/geom"

type GameDrawOptions struct {
	// Lerpable
	Time           float32
	Z              float32
	PlayerRotation geom.Vec3
	PlayerColor    geom.Vec3
	// Hard to predict
	PlayerPosition geom.Vec2
}
