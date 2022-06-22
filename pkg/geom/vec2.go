package geom

import "math"

type Vec2 struct {
	X, Y float32
}

func (v *Vec2) Add(v2 Vec2) {
	v.X += v2.X
	v.Y += v2.Y
}

func (v *Vec2) Sub(v2 Vec2) {
	v.X -= v2.X
	v.Y -= v2.Y
}

func (v *Vec2) MulN(n float32) {
	v.X *= n
	v.Y *= n
}

func (v *Vec2) Rotate(center Vec2, angle float32) {
	s, c := math.Sincos(float64(angle))
	v.Sub(center)
	x := v.X*float32(c) - v.Y*float32(s)
	y := v.X*float32(s) + v.Y*float32(c)
	v.X, v.Y = x, y
	v.Add(center)
}

func (v *Vec2) Dot(v2 Vec2) float32 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v *Vec2) Hypot(v2 Vec2) float32 {
	tv := *v
	tv.Sub(v2)

	return tv.Dot(tv)
}

func (v *Vec2) Project(v2 Vec2) Vec2 {
	k := v.Dot(v2) / v2.Dot(v2)

	return Vec2{
		X: k * v2.X,
		Y: k * v2.Y,
	}
}

func (v *Vec2) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v *Vec2) Normalize() {
	l := v.Length()
	if l > 0 {
		v.X /= l
		v.Y /= l
	}
}

func (v *Vec2) Atan2() float32 {
	return float32(math.Atan2(float64(v.Y), float64(v.X)))
}

func (v *Vec2) DistanceTo(v2 Vec2) float32 {
	x0 := float64(v.X)
	y0 := float64(v.Y)
	x1 := float64(v2.X)
	y1 := float64(v2.Y)
	return float32(math.Sqrt((x0-x1)*(x0-x1) + (y0-y1)*(y0-y1)))
}

func (v *Vec2) DistanceSqTo(v2 Vec2) float32 {
	return (v.X-v2.X)*(v.X-v2.X) + (v.Y-v2.Y)*(v.Y-v2.Y)
}

func (v *Vec2) Equals(v2 Vec2) bool {
	return v.X == v2.X && v.Y == v2.Y
}

func (v *Vec2) IsZero() bool {
	return v.X == 0 && v.Y == 0
}
