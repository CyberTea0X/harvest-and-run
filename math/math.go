package math

import (
	"golang.org/x/exp/constraints"
	"math"
)

func Distance[T constraints.Float | constraints.Integer](x1, x2, y1, y2 T) T {
	return T(math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))))
}
