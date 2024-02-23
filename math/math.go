package math

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func Distance[T Number](x1, y1, x2, y2 T) T {
	return T(math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))))
}

func Scale[T, A Number](slice [2]T, scale A) [2]T {
	slice[0] *= T(scale)
	slice[1] *= T(scale)
	return slice
}
