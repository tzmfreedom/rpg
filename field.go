package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

func fieldHandleInput() {
	player.Update()
}

type Player struct{
	x, y int
	speed, animation int
	vx, vy float64
}

func NewPlayer(x, y, speed int) *Player {
	return &Player{x: x, y: y, speed: speed}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// draw player
	img := ebiten.NewImage(size, size)
	img.Fill(color.RGBA{0x00, 0xff, 0x00, 0xff})
	options := &ebiten.DrawImageOptions{}
	if p.animation > 0 {
		if p.vx == 0 {
			diff := p.vy * float64(p.speed - p.animation)
			options.GeoM.Translate(float64(p.x * size), float64(p.y * size) + diff)
		} else {
			diff := p.vx * float64(p.speed - p.animation)
			options.GeoM.Translate(float64(p.x * size) + diff, float64(p.y * size))
		}
	} else {
		options.GeoM.Translate(float64(p.x * size), float64(p.y * size))
	}
	screen.DrawImage(img, options)
}

func (p *Player) Update() {
	if p.animation > 0 {
		p.animation--
		if p.animation == 0 {
			if p.vx > 0 {
				p.x++
			} else if p.vx < 0 {
				p.x--
			}
			if p.vy > 0 {
				p.y++
			} else if p.vy < 0 {
				p.y--
			}
			p.vx = 0
			p.vy = 0
		}
		return
	}
	velocity := float64(size)/float64(p.speed)
	if v := inpututil.KeyPressDuration(ebiten.KeyLeft); v > 0 {
		if p.x > 0 {
			if worldMap.Accessable(p.y, p.x-1) {
				p.vx = -velocity
				p.vy = 0
				p.animation = p.speed
			}
		}
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyRight); v > 0 {
		if p.x < worldMap.xmax {
			if worldMap.Accessable(p.y, p.x+1) {
				p.vx = velocity
				p.vy = 0
				p.animation = p.speed
			}
		}
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyDown); v > 0 {
		if p.y < worldMap.ymax {
			if worldMap.Accessable(p.y+1, p.x) {
				p.vx = 0
				p.vy = velocity
				p.animation = p.speed
			}
		}
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyUp); v > 0 {
		if p.y > 0 {
			if worldMap.Accessable(p.y-1, p.x) {
				p.vx = 0
				p.vy = -velocity
				p.animation = p.speed
			}
		}
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyW); v == 1 {
		phase = PhaseMenu
	}
	if v := inpututil.KeyPressDuration(ebiten.KeyS); v == 1 {
		if worldMap.HasEvent(p.x, p.y) {
			message = "hogehoge!"
		}
	}
}

type WorldMap struct{
	xmax int
	ymax int
	data [][]*Block
}

func NewWorldMap(xmax, ymax int) *WorldMap {
	data := make([][]*Block, ymax)
	for i, _ := range data {
		data[i] = make([]*Block, xmax)
		for j, _ := range data[i] {
			data[i][j] = &Block{
				Access: true,
				Event:  false,
			}
		}
	}
	data[5][10] = &Block{
		Access: false,
		Event:  false,
	}
	data[10][10] = &Block{
		Access: true,
		Event:  true,
	}
	return &WorldMap{
		data: data,
		xmax: xmax,
		ymax: ymax,
	}
}

func (m *WorldMap) Draw(screen *ebiten.Image) {
	for y, line := range m.data {
		for x, block := range line {
			if block != nil && !block.Access {
				img := ebiten.NewImage(size, size)
				img.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
				options := &ebiten.DrawImageOptions{}
				options.GeoM.Translate(float64(x * size), float64(y * size))
				screen.DrawImage(img, options)
			}
			if block != nil && block.Event {
				img := ebiten.NewImage(size, size)
				img.Fill(color.RGBA{0x00, 0x00, 0xff, 0xff})
				options := &ebiten.DrawImageOptions{}
				options.GeoM.Translate(float64(x * size), float64(y * size))
				screen.DrawImage(img, options)
			}
		}
	}
}

func (m *WorldMap) Accessable(x, y int) bool {
	block := worldMap.data[y][x]
	return block != nil && block.Access
}

func (m *WorldMap) HasEvent(x, y int) bool {
	block := worldMap.data[y][x]
	return block != nil && block.Event
}

func fieldDraw(screen *ebiten.Image) {
	// text.Draw(screen, string(score), scoreFont, scoreX, scoreY, color.White)
	player.Draw(screen)

	worldMap.Draw(screen)
}
