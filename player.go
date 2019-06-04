package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var playerSprite *ebiten.Image

func init() {
	playerSprite, _, _ = ebitenutil.NewImageFromFile("data/player.png", ebiten.FilterDefault)
}

type Player struct {
	x   float64
	y   float64
	dir float64
}

func (p *Player) Init() {
	p.x = 0
	p.y = 0
	p.dir = 1
}

func (p *Player) Update(input Input) {

	var speed float64 = 1

	if input.x != 0 {
		p.dir = float64(input.x)
	}

	p.x += float64(input.x) * speed
	p.y += float64(input.y) * speed

}

func (p *Player) Draw(screen *ebiten.Image) {
	o := ebiten.DrawImageOptions{}
	o.GeoM.Scale(p.dir, 1)
	o.GeoM.Translate(p.x, p.y)
	screen.DrawImage(playerSprite, &o)

}
