package unit

import (
	"fmt"
	"harvest-and-run/control"
	"harvest-and-run/errors"
	"harvest-and-run/math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Unit struct {
	Id               int
	Image            *ebiten.Image
	Name             string
	Position         math.Vec2
	orders           []*control.Order
	MaxSpeed         float32
	CurrentSpeed     float32
	LineAcceleration float32
	// Angle
}

func (u *Unit) Order(o *control.Order) {
	u.orders = append(u.orders, o)
}

func (u *Unit) CurrentOrder() (*control.Order, error) {
	if len(u.orders) == 0 {
		return nil, errors.ErrNoOrder
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

func (u *Unit) ProcessOrders(dt time.Duration) {
	o, err := u.CurrentOrder()
	if err != nil {
		return
	}
	switch o.Command {
	case control.CommandMove:
		pos1 := u.Position
		pos2 := o.Position
		if !u.CanMove() || math.Distance(pos1[0], pos1[1], pos2[0], pos2[1]) <= 1 {
			u.CurrentSpeed = 0
			u.FinishOrder()
			return
		}
		u.MoveTo(int(o.Position[0]), int(o.Position[1]), dt)
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

func (u *Unit) Update(dt time.Duration) {
	u.ProcessOrders(dt)
}
