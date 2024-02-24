package player

import "harvest-and-run/control"

type Player struct {
	Selection control.UnitSelection
}

func New() *Player {
	p := new(Player)
	return p
}
