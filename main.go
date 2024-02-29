package main

import (
	"harvest-and-run/assets"
	"harvest-and-run/game"
	"harvest-and-run/player"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// cfg, err := config.FromToml("config.toml")
	// if err != nil {
	// panic(err)
	// }
	// ishost := false
	// if os.Getenv("HOST") == "1" {
	// ishost = true
	// }
	// fmt.Println(ishost)
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Harvest-And-Run")
	ebiten.SetVsyncEnabled(false)
	player := player.New()
	assets := assets.New()
	game := game.New(assets, player)
	x, y := ebiten.WindowSize()
	game.SpawnDrone(300, 300)
	drone := game.SpawnDrone(x/2, y/2)
	player.Selection.Add(drone.Id)
	game.Player = player
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
