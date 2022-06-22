package geom

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Triangle struct {
	A, B, C Vec2
}

func NewTriangle(a, b, c Vec2) *Triangle {
	return &Triangle{
		A: a,
		B: b,
		C: c,
	}
}

func distancePointSegment(a, b, p Vec2) float32 {
	ap := p
	ap.Sub(a)
	ab := b
	ab.Sub(a)

	// Get closest point on triangle side
	closest := ap.Project(ab)
	closest.Add(a)

	ad := closest
	ad.Sub(a)

	var k float32
	if math.Abs(float64(ab.X)) > math.Abs(float64(ab.Y)) {
		k = ad.X / ab.X
	} else {
		k = ad.Y / ab.Y
	}

	if k <= 0 {
		return float32(math.Sqrt(float64(p.Hypot(a))))
	} else if k >= 1 {
		return float32(math.Sqrt(float64(p.Hypot(b))))
	}

	return float32(math.Sqrt(float64(p.Hypot(closest))))
}

func (t *Triangle) Rotate(center Vec2, angle float32) {
	t.A.Rotate(center, angle)
	t.B.Rotate(center, angle)
	t.C.Rotate(center, angle)
}

func (t *Triangle) Translate(x, y float32) {
	t.A.Add(Vec2{
		X: x,
		Y: y,
	})
	t.B.Add(Vec2{
		X: x,
		Y: y,
	})
	t.C.Add(Vec2{
		X: x,
		Y: y,
	})
}

func (t *Triangle) Scale(v float32) {
	center := Vec2{
		X: (t.A.X + t.B.X + t.C.X) / 3,
		Y: (t.A.Y + t.B.Y + t.C.Y) / 3,
	}

	va := center
	va.Sub(t.A)
	va.MulN(v)
	va.Add(center)

	vb := center
	vb.Sub(t.B)
	vb.MulN(v)
	vb.Add(center)

	vc := center
	vc.Sub(t.C)
	vc.MulN(v)
	vc.Add(center)

	t.A = va
	t.B = vb
	t.C = vc
}

func (t *Triangle) IntersectsCircle(p Vec2, radius float32) bool {
	rSq := radius * radius
	// Check if one vertices is contained by the circle
	dist0 := p.DistanceSqTo(t.A)
	if dist0 <= rSq {
		return true
	}
	dist1 := p.DistanceSqTo(t.B)
	if dist1 <= rSq {
		return true
	}
	dist2 := p.DistanceSqTo(t.C)
	if dist2 <= rSq {
		return true
	}

	// Check if circle center is contained by the triangle
	a := ((t.B.Y-t.C.Y)*(p.X-t.C.X) + (t.C.X-t.B.X)*(p.Y-t.C.Y)) / ((t.B.Y-t.C.Y)*(t.A.X-t.C.X) + (t.C.X-t.B.X)*(t.A.Y-t.C.Y))
	b := ((t.C.Y-t.A.Y)*(p.X-t.C.X) + (t.A.X-t.C.X)*(p.Y-t.C.Y)) / ((t.B.Y-t.C.Y)*(t.A.X-t.C.X) + (t.C.X-t.B.X)*(t.A.Y-t.C.Y))
	c := 1 - a - b

	if (0 <= a && a <= 1) && (0 <= b && b <= 1) && (0 <= c && c <= 1) {
		return true
	}

	// Check if one of the sides is intersecting circle
	if distancePointSegment(t.B, t.A, p) <= radius {
		return true
	}
	if distancePointSegment(t.C, t.B, p) <= radius {
		return true
	}
	if distancePointSegment(t.A, t.C, p) <= radius {
		return true
	}

	return false
}

func (t *Triangle) AppendVerticesIndices(vertices []ebiten.Vertex, indices []uint16, index int, cr, cg, cb, ca float32) ([]ebiten.Vertex, []uint16) {
	vertices = append(vertices, []ebiten.Vertex{
		{
			DstX:   t.A.X,
			DstY:   t.A.Y,
			ColorR: cr,
			ColorG: cg,
			ColorB: cb,
			ColorA: ca,
		},
		{
			DstX:   t.B.X,
			DstY:   t.B.Y,
			ColorR: cr,
			ColorG: cg,
			ColorB: cb,
			ColorA: ca,
		},
		{
			DstX:   t.C.X,
			DstY:   t.C.Y,
			ColorR: cr,
			ColorG: cg,
			ColorB: cb,
			ColorA: ca,
		},
	}...)
	indices = append(
		indices,
		uint16(index)*3+0,
		uint16(index)*3+1,
		uint16(index)*3+2,
	)

	return vertices, indices
}
