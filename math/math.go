package math

import "math"

type Position struct {
	X int
	Y int
}

func NewPosition(x int, y int) Position {
	return Position{x, y}
}

func Distance(pos1 Position, pos2 Position) float64 {
	x1 := float64(pos1.X)
	x2 := float64(pos2.X)
	y1 := float64(pos1.Y)
	y2 := float64(pos2.Y)
	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}
