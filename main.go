package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var screenWidth, screenHeight int
var size int
var player *Player
var worldMap *WorldMap
var message string
var phase int

type Block struct {
	Access bool
	Event  bool
}

const (
	PhaseField = iota
	PhaseMenu
	PhaseBattle
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
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("RPG")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	phaseHandleFunctions[phase]()
	ebitenutil.DebugPrint(screen, message)
	phaseDrawFunctions[phase](screen)
	if player.x + player.y == 10 {
		phase = PhaseBattle
	}
}

func init() {
	screenWidth = 640
	screenHeight = 480
	size = 32
	phase = PhaseField

	player = NewPlayer(0, 0, 12)
	worldMap = NewWorldMap(screenWidth / size, screenHeight / size)
}