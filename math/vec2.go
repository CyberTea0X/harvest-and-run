package math

import "math"

type Vec2 [2]float32

func Vec2From[T Number](x, y T) Vec2 {
	return Vec2{float32(x), float32(y)}
}

func (v Vec2) X() float32 {
	return v[0]
}

func (v Vec2) Y() float32 {
	return v[1]
}

func (v Vec2) X64() float64 {
	return float64(v[0])
}

func (v Vec2) Y64() float64 {
	return float64(v[1])
}

func (v Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{v[0] - v2[0], v[1] - v2[1]}
}

func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v[0] + v2[0], v[1] + v2[1]}
}

func (v Vec2) Mul(v2 Vec2) Vec2 {
	return Vec2{v[0] * v2[0], v[1] * v2[1]}
}

func (v Vec2) Div(v2 Vec2) Vec2 {
	return Vec2{v[0] / v2[0], v[1] / v2[1]}
}

func (v Vec2) Clamp(min, max Vec2) Vec2 {
	for i := range v {
		if v[i] < min[i] {
			v[i] = min[i]
		} else if v[i] > max[i] {
			v[i] = max[i]
		}
	}
	return v
}

// Squared lenght of the vector
func (v Vec2) LengthSqr() float32 {
	return v[0]*v[0] + v[1]*v[1]
}

func (v Vec2) Scale(scale float32) Vec2 {
	v[0] *= scale
	v[1] *= scale
	return v
}

func (v Vec2) Normalize() Vec2 {
	sl := v.LengthSqr()
	if sl == 0 || sl == 1 {
		return v
	}
	return v.Scale(float32(1 / math.Sqrt(float64(sl))))
}
