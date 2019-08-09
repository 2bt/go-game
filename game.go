package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	world   *World
	hero    *Hero
	bullets Entities
}

func NewGame() (*Game, error) {
	g := &Game{
		world: NewWorld("data/level-1.txt"),
		hero: &Hero{
			x:   ScreenWidth / 2,
			y:   ScreenHeight / 2,
			dir: 1,
		},
	}
	return g, nil
}

func (g *Game) Update() error {
	g.hero.Update(getInput())
	g.bullets.Update()

	return nil
}

func (g *Game) AddBullet(b *Bullet) {
	g.bullets = append(g.bullets, b)
}

func (g *Game) Draw(screen *ebiten.Image) {

	cam := Box{
		g.hero.x - ScreenWidth/2,
		g.hero.y - ScreenHeight/2,
		ScreenWidth,
		ScreenHeight,
	}

	g.world.Draw(screen, &cam)
	g.hero.Draw(screen, &cam)
	g.bullets.Draw(screen, &cam)
	// ebitenutil.DebugPrint(screen, "Hello, World!")
}

var game *Game
