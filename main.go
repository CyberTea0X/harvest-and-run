package main

import (
	"fmt"
	"harvest-and-run/math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Harvest-And-Run")
	player := NewPlayer()
	game := new(Game)
	drone := new(Unit)
	img, _, err := ebitenutil.NewImageFromFile("./images/flying_bot.png")
	if err != nil {
		panic(err)
	}
	drone.Image = img
	drone.MaxSpeed = 2.0
	drone.LineAcceleration = 2.0
	x, y := ebiten.WindowSize()
	fmt.Println(x, y)
	x = x / 2
	y = y / 2
	fmt.Println(x, y)
	drone.Position = math.Position{x, y}
	droneId := game.AddUnit(drone)
	player.Selection.Add(droneId)
	game.Player = player
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
