package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image/color"
)

func fieldHandleInput() {
	if v := inpututil.KeyPressDuration(ebiten.KeyLeft); v == 1 || (v > 0 && v%speed == 0) {
		if x > 0 {
			if block := worldMap[y][x-1]; block != nil && block.Access {
				x--
			}
		}
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyRight); v == 1 || (v > 0 && v%speed == 0) {
		if x < xmax {
			if block := worldMap[y][x+1]; block != nil && block.Access {
				x++
			}
		}
	}
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
		phase = PHASE_MENU
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyS); v == 1 {
		if block := worldMap[y][x]; block != nil && block.Event {
			message = "hogehoge!"
		}
	}
}

func fieldDraw(screen *ebiten.Image) {
	img, _ := ebiten.NewImage(size, size, 0)
	img.Fill(color.RGBA{0x00, 0xff, 0x00, 0xff})
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x * size), float64(y * size))
	screen.DrawImage(img, options)
	// text.Draw(screen, string(score), scoreFont, scoreX, scoreY, color.White)

	for y, line := range worldMap {
		for x, block := range line {
			if block != nil && !block.Access {
				img, _ := ebiten.NewImage(size, size, 0)
				img.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
				options := &ebiten.DrawImageOptions{}
				options.GeoM.Translate(float64(x * size), float64(y * size))
				screen.DrawImage(img, options)
			}
			if block != nil && block.Event {
				img, _ := ebiten.NewImage(size, size, 0)
				img.Fill(color.RGBA{0x00, 0x00, 0xff, 0xff})
				options := &ebiten.DrawImageOptions{}
				options.GeoM.Translate(float64(x * size), float64(y * size))
				screen.DrawImage(img, options)
			}
		}
	}
}
