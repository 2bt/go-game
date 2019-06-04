package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var sprite *ebiten.Image

var x float64
var y float64

func init() {
	var _ error
	sprite, _, _ = ebitenutil.NewImageFromFile("sprite.png", ebiten.FilterDefault)
}

func update(screen *ebiten.Image) error {

	// toggle fullscreen
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x += 1
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	o := ebiten.DrawImageOptions{}
	o.GeoM.Translate(x, y)
	screen.DrawImage(sprite, &o)

	ebitenutil.DebugPrint(screen, "Hello, World!")

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
