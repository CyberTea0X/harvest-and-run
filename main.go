package main

import (
	"harvest-and-run/assets"
	"harvest-and-run/game"
	"harvest-and-run/player"
	"harvest-and-run/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Harvest-And-Run")
	player := player.New()
	game := new(game.Game)
	game.Assets = assets.SetupDefault()
	font1, err := game.Assets.Font("Timeburner.ttf")
	if err != nil {
		panic(err)
	}
	game.Ui = ui.Setup(font1)
	x, y := ebiten.WindowSize()
	game.SpawnDrone(300, 300)
	drone := game.SpawnDrone(x/2, y/2)
	player.Selection.Add(drone.Id)
	game.Player = player
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
