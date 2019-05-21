package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image/color"
	"log"
)

var x, y, xmax, ymax int
var screenWidth, screenHeight int
var size int
var speed int
var moving bool
var worldMap [][]*Block
var message string
var phase int

type Block struct {
	Access bool
	Event bool
}

const (
	PHASE_FIELD = iota
	PHASE_MENU
	PHASE_BATTLE
)

var phaseHandleFunctions = []func(){
	fieldHandleInput,
	menuHandleInput,
	battleHandleInput,
}

var phaseDrawFunctions = []func(*ebiten.Image){
	fieldDraw,
	menuDraw,
	battleDraw,
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "RPG"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	phaseHandleFunctions[phase]()
	ebitenutil.DebugPrint(screen, message)
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	phaseDrawFunctions[phase](screen)
	if x + y == 10 {
		phase = PHASE_BATTLE
	}
	return nil
}

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

func init() {
	screenWidth = 640
	screenHeight = 480
	size = 32
	xmax = screenWidth / size
	ymax = screenHeight / size
	speed = 12
	worldMap = make([][]*Block, ymax)
	for i, _ := range worldMap {
		worldMap[i] = make([]*Block, xmax)
		for j, _ := range worldMap[i] {
			worldMap[i][j] = &Block{
				Access: true,
				Event: false,
			}
		}
	}
	phase = PHASE_FIELD
	worldMap[5][10] = &Block{
		Access: false,
		Event: false,
	}
	worldMap[10][10] = &Block{
		Access: true,
		Event: true,
	}
}