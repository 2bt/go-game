package main

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var playerSprite *ebiten.Image

func init() {
	playerSprite, _, _ = ebitenutil.NewImageFromFile("data/player.png", ebiten.FilterDefault)
}

type Player struct {
	x    float64
	y    float64
	vx   float64
	dir  float64
	tick int
}

func (p *Player) Update(input Input) {

	var speed float64 = 1

	if input.x != 0 {
		p.dir = float64(input.x)
	}

	p.vx = float64(input.x) * speed
	p.x += p.vx
	p.y += float64(input.y) * speed

	p.tick++
}

func (p *Player) Draw(screen *ebiten.Image) {
	o := ebiten.DrawImageOptions{}
	if p.dir == -1 {
		o.GeoM.Scale(p.dir, 1)
		o.GeoM.Translate(16, 0)
	}
	o.GeoM.Translate(p.x, p.y)

	f := 0
	if p.vx != 0 {
		f = p.tick/4%2 + 1
	}

	frame := playerSprite.SubImage(image.Rect(f*16, 0, f*16+16, 16)).(*ebiten.Image)

	screen.DrawImage(frame, &o)

}
