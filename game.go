package main

import (
	"harvest-and-run/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	RemotePlayers []*Player
	Player        *Player
	MaxUnitId     int
	units         []*Unit
}

// Adds new unit to the game.
//
// returns unit unique identifier
func (g *Game) AddUnit(unit *Unit) int {
	g.units = append(g.units, unit)
	return len(g.units) - 1
}

func (g *Game) GetUnit(id int) *Unit {
	return g.units[id]
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		og := []*Order{}
		for _, uId := range g.Player.Selection.Units {
			u := g.GetUnit(uId)
			o := new(Order)
			og = append(og, o)
			u.Order(o)
			o.OrderGroup = og
			o.Command = CommandMove
			o.SourceUnit = uId
			o.Position = math.NewPosition(ebiten.CursorPosition())
		}
	}
	for _, u := range g.units {
		u.Update(g)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, unit := range g.units {
		unit.Draw(screen)
	}
}

func (g *Game) ProcessCommands(p *Player) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}
