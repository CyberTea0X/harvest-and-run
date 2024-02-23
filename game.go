package main

import (
	"harvest-and-run/math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	RemotePlayers []*Player
	Player        *Player
	MaxUnitId     int
	units         []*Unit
	lastUpdate    time.Time
}

// Returns amount of time passed since last update
func (g *Game) Dt() time.Duration {
	dt := time.Since(g.lastUpdate)
	if dt.Milliseconds() == 0 {
		return time.Duration(time.Millisecond * time.Duration(1000/ebiten.TPS()))
	}
	return dt
}

// Adds new unit to the game.
//
// Generates and sets unit unique ID
func (g *Game) AddUnit(u *Unit) {
	g.units = append(g.units, u)
	u.Id = len(g.units) - 1
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
			o.Position = math.Vec2From(ebiten.CursorPosition())
		}
	}
	for _, u := range g.units {
		u.Update(g)
	}
	g.lastUpdate = time.Now()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, unit := range g.units {
		unit.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}
