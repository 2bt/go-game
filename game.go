package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	world     World
	hero      *Hero
	bullets   Entities
	mobs      Entities
	particles Entities
}

func NewGame() (*Game, error) {
	var g Game
	g.world.Load("data/level-1.txt", func(t byte, x, y float64) {
		switch t {
		case '@':
			g.hero = NewHero(x, y)
		case 'M':
			g.mobs = append(g.mobs, NewMob(x, y))
		}
	})
	return &g, nil
}

func (g *Game) Update() {
	g.hero.Update(getInput())
	g.bullets.Update()
	g.particles.Update()
	g.mobs.Update()
}

func (g *Game) AddBullet(b *Bullet) {
	g.bullets = append(g.bullets, b)
}

func (g *Game) Draw(screen *ebiten.Image) {

	cam := &Box{
		g.hero.x - ScreenWidth/2,
		g.hero.y - ScreenHeight/2 - 30,
		ScreenWidth,
		ScreenHeight,
	}

	g.world.Draw(screen, cam)
	g.particles.Draw(screen, cam)
	g.bullets.Draw(screen, cam)
	g.hero.Draw(screen, cam)
	g.mobs.Draw(screen, cam)
}

var game *Game
