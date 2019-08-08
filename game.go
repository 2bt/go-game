package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

type Game struct {
	world   *World
	hero    *Hero
	bullets []*Bullet
}

func NewGame() (*Game, error) {
	g := &Game{
		world: NewWorld("data/level-1.txt"),
		hero: &Hero{
			x:   screenWidth / 2,
			y:   screenHeight / 2,
			dir: 1,
		},
	}
	return g, nil
}

func (g *Game) Update() error {
	g.hero.Update(getInput())
	for _, b := range g.bullets {
		b.Update()
	}
	return nil
}

func (g *Game) AddBullet(x float64, y float64, dir Dir) {
	b := &Bullet{
		x,
		y,
		dir,
	}

	g.bullets = append(g.bullets, b)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	g.hero.Draw(screen)
	for _, b := range g.bullets {
		b.Draw(screen)
	}
	// ebitenutil.DebugPrint(screen, "Hello, World!")
}

var game *Game

type Bullet struct {
	x   float64
	y   float64
	dir Dir
}

func (b *Bullet) Update() {
	// set new pos
	if b.dir == Right {
		b.x += 2.5
	} else {
		b.x -= 2.5
	}
}

func (h *Bullet) Draw(screen *ebiten.Image) {
	o := ebiten.DrawImageOptions{}
	o.GeoM.Translate(-16, -24)
	if h.dir == Left {
		o.GeoM.Scale(-1, 1)
	}
	o.GeoM.Translate(h.x, h.y)
	ebitenutil.DrawRect(screen, h.x-7, h.y-19, 14, 19, color.RGBA{100, 0, 0, 100})
}
