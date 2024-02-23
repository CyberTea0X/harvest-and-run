package main

import (
	"fmt"
	"harvest-and-run/math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Unit struct {
	Id               int
	Image            *ebiten.Image
	Name             string
	Position         math.Vec2
	orders           []*Order
	MaxSpeed         float32
	CurrentSpeed     float32
	LineAcceleration float32
	// Angle
}

func (u *Unit) Order(o *Order) {
	u.orders = append(u.orders, o)
}

func (u *Unit) CurrentOrder() (*Order, error) {
	if len(u.orders) == 0 {
		return nil, ErrNoOrder
	}
	return u.orders[0], nil
}

// Marks order as completed and then removes it from the orders
func (u *Unit) FinishOrder() {
	o := u.orders[len(u.orders)-1]
	o.Completed = true
	u.orders = u.orders[1:]
}

func (u *Unit) CancelOrder() {
	if len(u.orders) > 0 {
		u.orders = u.orders[1:]
	}
}

func (u *Unit) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	x := u.Position.X64() - float64(u.Image.Bounds().Dx()/2)
	y := u.Position.Y64() - float64(u.Image.Bounds().Dy()/2)
	op.GeoM.Translate(x, y)
	screen.DrawImage(u.Image, op)
}

func (u *Unit) CanMove() bool {
	return u.MaxSpeed > 0 && u.LineAcceleration > 0
}

func (u *Unit) Accelerate() {
	if u.CurrentSpeed < u.MaxSpeed {
		u.CurrentSpeed += u.LineAcceleration
		if u.CurrentSpeed > u.MaxSpeed {
			u.CurrentSpeed = u.MaxSpeed
		}
	}
}

func (u *Unit) MoveTo(x int, y int, dt time.Duration) {
	u.Accelerate()
	target := math.Vec2From(x, y)
	dir := target.Sub(u.Position).Normalize()
	nextPos := math.Vec2From(u.Position[0], u.Position[1])
	fmt.Println(dt.Milliseconds())
	nextPos[0] += dir[0] * u.CurrentSpeed * float32(dt.Milliseconds())
	nextPos[1] += dir[1] * u.CurrentSpeed * float32(dt.Milliseconds())
	dBefore := math.Distance(u.Position[0], u.Position[1], target[0], target[1])
	dAfter := math.Distance(nextPos[0], nextPos[1], target[0], target[1])
	if dAfter >= dBefore {
		nextPos = target
	}
	u.Position = nextPos
}

func (u *Unit) ProcessOrders(g *Game) {
	o, err := u.CurrentOrder()
	if err != nil {
		return
	}
	switch o.Command {
	case CommandMove:
		pos1 := u.Position
		pos2 := o.Position
		if !u.CanMove() || math.Distance(pos1[0], pos1[1], pos2[0], pos2[1]) <= 1 {
			u.CurrentSpeed = 0
			u.FinishOrder()
			return
		}
		u.MoveTo(int(o.Position[0]), int(o.Position[1]), g.Dt())
	}
}

func (u *Unit) Update(g *Game) {
	u.ProcessOrders(g)
}

// Creates Drone Unit
func NewDrone(x, y int) *Unit {
	drone := new(Unit)
	drone.Position = math.Vec2From(x, y)
	img, _, err := ebitenutil.NewImageFromFile("./images/flying_bot.png")
	if err != nil {
		panic(err)
	}
	drone.Image = img
	drone.MaxSpeed = 2.0
	drone.LineAcceleration = 0.1
	return drone
}
