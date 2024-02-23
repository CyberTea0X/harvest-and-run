package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Harvest-And-Run")
	player := NewPlayer()
	game := new(Game)
	x, y := ebiten.WindowSize()
	drone1 := NewDrone(300, 300)
	game.AddUnit(drone1)
	drone2 := NewDrone(x/2, y/2)
	game.AddUnit(drone2)
	player.Selection.Add(drone2.Id)
	game.Player = player
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
