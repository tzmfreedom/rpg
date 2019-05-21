package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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