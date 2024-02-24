package ui

import (
	"github.com/ebitenui/ebitenui"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameUI struct {
	ui    *ebitenui.UI
	font1 *truetype.Font
}

func Setup(font1 *truetype.Font) *GameUI {
	ui := new(GameUI)
	ui.ui = &ebitenui.UI{}
	return ui
}

func (g *GameUI) Update() {
	if g.ui.Container == nil {
		return
	}
	g.ui.Update()
}

func (g *GameUI) Draw(screen *ebiten.Image) {
	if g.ui.Container == nil {
		return
	}
	g.ui.Draw(screen)
}
