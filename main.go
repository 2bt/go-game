package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	screenWidth  = 400
	screenHeight = 225
)

type Game struct {
	world *World
	hero  *Hero
}

func NewGame() (*Game, error) {
	g := &Game{
		world: NewWorld("data/level-1.txt"),
		hero: &Hero{
			x:   screenWidth/2 - 8,
			y:   screenHeight/2 - 8,
			dir: 1,
		},
	}
	return g, nil
}

type Input struct {
	x int
	y int
	a bool
	b bool
}

func getInput() Input {
	var input Input
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		input.y--
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		input.y++
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		input.x--
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		input.x++
	}
	input.a = ebiten.IsKeyPressed(ebiten.KeyX)
	input.b = ebiten.IsKeyPressed(ebiten.KeyC)

	return input
}

func (g *Game) Update() error {
	g.hero.Update(getInput())
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	g.hero.Draw(screen)
	// ebitenutil.DebugPrint(screen, "Hello, World!")
}

var game *Game

func update(screen *ebiten.Image) error {
	if err := game.Update(); err != nil {
		return err
	}

	// toggle fullscreen
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	game.Draw(screen)
	return nil
}

func main() {
	var err error
	game, err = NewGame()
	if err != nil {
		log.Fatal(err)
	}
	err = ebiten.Run(update, screenWidth, screenHeight, 3, "go-game")
	if err != nil {
		log.Fatal(err)
	}
}
