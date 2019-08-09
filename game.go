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
			x:   screenWidth / 2,
			y:   screenHeight / 2,
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
	g.world.Draw(screen)
	g.hero.Draw(screen)
	g.bullets.Draw(screen)
	// ebitenutil.DebugPrint(screen, "Hello, World!")
}

var game *Game
