package main

import (
	"github.com/hajimehoshi/ebiten"
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

	i := 0
	for _, b := range g.bullets {
		if b.Update() {
			g.bullets[i] = b
			i++
		}
	}
	g.bullets = g.bullets[:i]

	return nil
}

func (g *Game) AddBullet(b *Bullet) {
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
