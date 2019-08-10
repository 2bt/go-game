package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

type Life struct {
	hp uint
	x  float64
	y  float64
}

func (h *Life) Draw(screen *ebiten.Image, cam *Box) {
	ebitenutil.DrawRect(screen, h.x-4-cam.X, h.y-24-cam.Y, 8, 2, color.RGBA{255, 0, 0, 200})
}
