package main

type Player struct {
	Selection UnitSelection
}

func NewPlayer() *Player {
	p := new(Player)
	return p
}
