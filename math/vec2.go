package math

type Vec2 [2]float32

func Vec2From(x float32, y float32) Vec2 {
	return Vec2{x, y}
}

func Vec2From64(x float64, y float64) Vec2 {
	return Vec2{float32(x), float32(y)}
}

func (v *Vec2) X() float32 {
	return v[0]
}

func (v *Vec2) Y() float32 {
	return v[1]
}

func (v *Vec2) X64() float64 {
	return float64(v[0])
}

func (v *Vec2) Y64() float64 {
	return float64(v[1])
}
