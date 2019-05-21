package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var battleCursor int

func battleHandleInput() {
	if v := inpututil.KeyPressDuration(ebiten.KeyDown); v == 1 || (v > 0 && v%speed == 0) {
		if y < ymax {
			if block := worldMap[y+1][x]; block != nil && block.Access {
				y++
			}
		}
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyUp); v == 1 || (v > 0 && v%speed == 0) {
		if y > 0 {
			if block := worldMap[y-1][x]; block != nil && block.Access {
				y--
			}
		}
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyS); v == 1 {
		x++
		phase = PHASE_FIELD
	}
}

func battleDraw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "battle!")
}
