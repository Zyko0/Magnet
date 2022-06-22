package core

import (
	"math"
	"math/rand"

	"github.com/Zyko0/Magnet/assets"
	"github.com/Zyko0/Magnet/logic"
	"github.com/Zyko0/Magnet/pkg/geom"
)

var ()

type Obstacle struct {
	rotationSpeed float32
	triangleIndex int

	Z         float32
	Angle     float32
	Triangles []*geom.Triangle
}

func newObstacle(currentDepth float32) *Obstacle {
	const (
		zAdd             = 10.
		maxRotationSpeed = 0.05
	)

	// TODO: random index generation from a local slice based on difficulty
	index := assets.ShapeIndexTrianglePlus
	triangles := make([]*geom.Triangle, len(assets.TriangleShapes[index]))
	for i, t := range assets.TriangleShapes[index] {
		tri := *t
		triangles[i] = &tri
	}

	return &Obstacle{
		rotationSpeed: maxRotationSpeed, // rand.Float32() * maxRotationSpeed,
		triangleIndex: index,

		Z:         currentDepth + zAdd,
		Angle:     rand.Float32() * 2 * math.Pi,
		Triangles: triangles,
	}
}

func (o *Obstacle) Update() {
	o.Angle += o.rotationSpeed
	if o.Angle > 2*math.Pi {
		o.Angle -= 2 * math.Pi
	}

	for i, t := range o.Triangles {
		*t = *assets.TriangleShapes[o.triangleIndex][i]
		t.Rotate(logic.Center, o.Angle)
		// TODO: scale
	}
}
