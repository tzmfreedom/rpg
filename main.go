package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image/color"
	"log"
)

var x, y, xmax, ymax int
var screenWidth, screenHeight int
var size int

func update(screen *ebiten.Image) error {
	handleInput()
	// ebitenutil.DebugPrint(screen, strconv.Itoa(score))
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	draw(screen)
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "RPG"); err != nil {
		log.Fatal(err)
	}
}

func handleInput() {
	if v := inpututil.KeyPressDuration(ebiten.KeyLeft); v == 1 || (v > 0 && v%30 == 0) {
		if x > 0 {
			x--
		}
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyRight); v == 1 || (v > 0 && v%30 == 0) {
		if x < xmax {
			x++
		}
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyDown); v == 1 || (v > 0 && v%30 == 0) {
		if y < ymax {
			y++
		}
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyUp); v == 1 || (v > 0 && v%30 == 0) {
		if y > 0 {
			y--
		}
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyS); v == 1 || (v > 0 && v%10 == 0) {
	}
}

func draw(screen *ebiten.Image) {
	img, _ := ebiten.NewImage(size, size, 0)
	img.Fill(color.RGBA{0x00, 0xff, 0x00, 0xff})
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, options)
	// text.Draw(screen, string(score), scoreFont, scoreX, scoreY, color.White)
}

func init() {
	screenWidth = 640
	screenHeight = 480
	size = 32
	xmax = screenWidth / size
	ymax = screenHeight / size
}