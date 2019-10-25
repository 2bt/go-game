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
		case 'm':
			g.AddMob(NewMob(x, y))
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

func (g *Game) AddParticle(e Entity) {
	g.particles = append(g.particles, e)
}

func (g *Game) AddBullet(e Entity) {
	g.bullets = append(g.bullets, e)
}

func (g *Game) AddMob(e Entity) {
	g.mobs = append(g.mobs, e)
}

func (g *Game) Draw(screen *ebiten.Image) {

	cam := Box{
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
