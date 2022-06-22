package geom

func Min(x, y float32) float32 {
	if x < y {
		return x
	}
	return y
}

func Max(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}

func Clamp(x, min, max float32) float32 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
