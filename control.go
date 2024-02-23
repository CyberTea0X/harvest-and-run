package main

import "harvest-and-run/math"

const (
	CommandMove int = iota
)

// Some unit order
type Order struct {
	// Other orders of same command and ordered at the same time
	OrderGroup []*Order
	// What to do
	Command int
	// Where to go, where to build, where to attack etc
	Position math.Vec2
	// source unit id
	SourceUnit int
	Completed  bool
}

type UnitSelection struct {
	Units []int
}

func (s *UnitSelection) Add(unitId int) {
	s.Units = append(s.Units, unitId)
}
