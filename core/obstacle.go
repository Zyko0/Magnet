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

	Z            float32
	Angle        float32
	Scale        float32
	Triangles    []*geom.Triangle
	SrcTriangles []*geom.Triangle
}

func newObstacle(z float32) *Obstacle {
	const (
		maxRotationSpeed = 0.025
	)

	// TODO: random index generation from a local slice based on difficulty
	index := assets.ShapeIndexTrianglePlus
	triangles := make([]*geom.Triangle, len(assets.TriangleShapes[index]))
	for i, t := range assets.TriangleShapes[index] {
		tri := *t
		triangles[i] = &tri
	}
	srcTriangles := make([]*geom.Triangle, len(assets.TriangleShapes[index]))
	for i, t := range assets.TriangleShapes[index] {
		tri := *t
		srcTriangles[i] = &tri
	}

	var sign float32 = 1
	if rand.Intn(2) == 0 {
		sign = -1
	}

	return &Obstacle{
		rotationSpeed: rand.Float32() * maxRotationSpeed * sign,
		triangleIndex: index,

		Z:            z,
		Angle:        rand.Float32() * 2 * math.Pi,
		Scale:        0,
		Triangles:    triangles,
		SrcTriangles: srcTriangles,
	}
}

func (o *Obstacle) Update(z float32) {
	// Note: sqrt(3.6) => why, idk
	const scaleFactor = 1.8973

	o.Angle += o.rotationSpeed // TODO: uncomment
	if o.Angle > 2*math.Pi {
		o.Angle -= 2 * math.Pi
	}

	// Note: 0.5 is another magic number :)
	o.Scale = (o.Z - z + 0.5) * scaleFactor
	o.Scale /= (o.Scale * o.Scale)
	o.Scale = geom.Clamp(o.Scale, 0, 1)
	for i, t := range o.Triangles {
		*t = *assets.TriangleShapes[o.triangleIndex][i]
		t.Rotate(logic.Center, o.Angle)
		*o.SrcTriangles[i] = *t
		t.Scale(logic.Center, o.Scale)
	}
}
