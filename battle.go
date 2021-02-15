package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var battleCursor int

func battleHandleInput() {
	if v := inpututil.KeyPressDuration(ebiten.KeyDown); v > 0 {
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyUp); v > 0 {
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyS); v > 0 {
		player.x++
		phase = PhaseField
	}
}

func battleDraw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "battle!")
}
