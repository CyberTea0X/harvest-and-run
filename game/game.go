package game

import (
	"harvest-and-run/assets"
	"harvest-and-run/control"
	"harvest-and-run/math"
	"harvest-and-run/player"
	"harvest-and-run/ui"
	"harvest-and-run/unit"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Assets        *assets.Assets
	RemotePlayers []*player.Player
	Player        *player.Player
	MaxUnitId     int
	units         []*unit.Unit
	lastUpdate    time.Time
	Ui            *ui.GameUI
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
func (g *Game) AddUnit(u *unit.Unit) {
	g.units = append(g.units, u)
	u.Id = len(g.units) - 1
}

func (g *Game) GetUnit(id int) *unit.Unit {
	return g.units[id]
}

func (g *Game) Update() error {
	g.Ui.Update()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		og := []*control.Order{}
		for _, uId := range g.Player.Selection.Units {
			u := g.GetUnit(uId)
			o := new(control.Order)
			og = append(og, o)
			u.Order(o)
			o.OrderGroup = og
			o.Command = control.CommandMove
			o.SourceUnit = uId
			o.Position = math.Vec2From(ebiten.CursorPosition())
		}
	}
	for _, u := range g.units {
		u.Update(g.Dt())
	}
	g.lastUpdate = time.Now()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Ui.Draw(screen)
	for _, unit := range g.units {
		unit.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

// Creates Drone Unit
func (g *Game) SpawnDrone(x, y int) *unit.Unit {
	drone := new(unit.Unit)
	drone.Position = math.Vec2From(x, y)
	img, err := g.Assets.Image("flying_bot.png")
	if err != nil {
		panic(err)
	}
	drone.Image = img
	drone.MaxSpeed = 2.0
	drone.LineAcceleration = 0.1
	g.AddUnit(drone)
	return drone
}
