package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var menuCursor int

func menuHandleInput() {
	if v := inpututil.KeyPressDuration(ebiten.KeyDown); v > 0 {
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyUp); v > 0 {
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyW); v == 1 {
		phase = PhaseField
	}
}

func menuDraw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "menu!")
}
