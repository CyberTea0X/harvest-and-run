package math

import "fmt"

type Position [2]int

func NewPosition(x int, y int) Position {
	fmt.Println(x, y)
	return Position{x, y}
}
