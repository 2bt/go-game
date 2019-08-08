package main

import "github.com/hajimehoshi/ebiten"

type Input struct {
	x     int
	y     int
	jump  bool
	shoot bool
}

//todo
// make keys configurable during runtime
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
	input.jump = ebiten.IsKeyPressed(ebiten.KeyX) || ebiten.IsKeyPressed(ebiten.KeySpace)
	input.shoot = ebiten.IsKeyPressed(ebiten.KeyW)

	return input
}
