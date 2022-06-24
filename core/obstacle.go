package core

import (
	"math"
	"math/rand"

	"github.com/Zyko0/Magnet/assets"
	"github.com/Zyko0/Magnet/logic"
	"github.com/Zyko0/Magnet/pkg/geom"
)

const (
	ObstacleKindDeath byte = iota
	ObstacleKindPortalNone
	ObstacleKindPortalNorth
	ObstacleKindPortalSouth
)

type Obstacle struct {
	creationTick  uint64
	kind          byte
	rotationSpeed float32
	triangleIndex int

	Z            float32
	Angle        float32
	Scale        float32
	Triangles    []*geom.Triangle
	SrcTriangles []*geom.Triangle
}

func newObstacle(ticks uint64, z float32, kind byte) *Obstacle {
	const (
		minRotationSpeed = 0.0125
		maxRotationSpeed = 0.025
	)

	// TODO: random index generation from a local slice based on difficulty
	index := 1 + rand.Intn(assets.ShapeIndexEmptyCross)
	// If it's a portal set the index to the quad shape
	if kind != ObstacleKindDeath {
		index = assets.ShapeIndexPortal
	}
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
		creationTick:  ticks,
		kind:          kind,
		rotationSpeed: sign * (minRotationSpeed + rand.Float32()*(maxRotationSpeed-minRotationSpeed)),
		triangleIndex: index,

		Z:            z,
		Angle:        rand.Float32() * 2 * math.Pi,
		Scale:        0,
		Triangles:    triangles,
		SrcTriangles: srcTriangles,
	}
}

var colorWhite = []float32{1, 1, 1}

func (o *Obstacle) GetColor() []float32 {
	switch o.kind {
	case ObstacleKindDeath:
		return colorWhite
	case ObstacleKindPortalNone:
		return colorsByAttraction[AttractionNone]
	case ObstacleKindPortalNorth:
		return colorsByAttraction[AttractionNorth]
	case ObstacleKindPortalSouth:
		return colorsByAttraction[AttractionSouth]
	}

	return colorWhite
}

func (o *Obstacle) GetAlpha() float32 {
	if o.kind == ObstacleKindDeath {
		return 0.9
	}
	return 0.5
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
		*o.SrcTriangles[i] = *assets.TriangleShapes[o.triangleIndex][i]
		*t = *assets.TriangleShapes[o.triangleIndex][i]
		t.Rotate(logic.Center, o.Angle)
		t.Scale(logic.Center, o.Scale)
	}
}
