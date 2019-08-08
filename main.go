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
