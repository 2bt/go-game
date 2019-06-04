package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var player Player

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
	return input
}

func update(screen *ebiten.Image) error {

	// toggle fullscreen
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	player.Update(getInput())

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	player.Draw(screen)

	// ebitenutil.DebugPrint(screen, "Hello, World!")

	return nil
}

func main() {

	player.Init()

	err := ebiten.Run(update, screenWidth, screenHeight, 2, "go-game")
	if err != nil {
		log.Fatal(err)
	}
}
