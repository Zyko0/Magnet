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
