package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var menuCursor int

func menuHandleInput() {
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
	if v := inpututil.KeyPressDuration(ebiten.KeyW); v == 1 {
		phase = PHASE_FIELD
	}
}

func menuDraw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "menu!")
}
