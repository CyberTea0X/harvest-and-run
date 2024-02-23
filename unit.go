package main

import "github.com/hajimehoshi/ebiten/v2"

type Unit struct {
	Image            *ebiten.Image
	Name             string
	Position         Position
	orders           []*Order
	MaxSpeed         int
	CurrentSpeed     float64
	LineAcceleration int
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
	x := float64(u.Position.X - u.Image.Bounds().Dx()/2)
	y := float64(u.Position.Y - u.Image.Bounds().Dy()/2)
	op.GeoM.Translate(x, y)
	screen.DrawImage(u.Image, op)
}

func (u *Unit) CanMove() bool {
	return u.MaxSpeed > 0 && u.LineAcceleration > 0
}

func (u *Unit) Move(pos Position) {
	u.CurrentSpeed += float64(u.LineAcceleration)
	u.Position = pos
}

func (u *Unit) ProcessOrders(g *Game) {
	o, err := u.CurrentOrder()
	if err != nil {
		return
	}
	switch o.Command {
	case CommandMove:
		if !u.CanMove() || Distance(u.Position, o.Position) <= 1 {
			u.CurrentSpeed = 0
			u.FinishOrder()
			return
		}
		u.Move(o.Position)
	}
}

func (u *Unit) Update(g *Game) {
	u.ProcessOrders(g)
}
