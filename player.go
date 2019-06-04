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
	x float64
	y float64
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.x -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.x += 1
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	o := ebiten.DrawImageOptions{}
	o.GeoM.Translate(p.x, p.y)
	screen.DrawImage(playerSprite, &o)

}
