package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

type Life struct {
	hp        uint
	x         float64
	y         float64
	ownerSize float64
	damage    uint
	width     float64
	alive     bool
	xOffset   float64
	yOffset   float64
}

func (h *Life) Update(x, y float64) bool {
	h.x = x - h.xOffset
	h.y = y - h.yOffset

	h.width = 10

	if h.ownerSize > 0 {
		h.width = h.ownerSize / 3
	}

	if h.damage > 0 && h.hp > 0 {
		//reduce the health bar width by the same percentage the health was reduced
		h.width -= float64(h.damage) / float64(h.hp) * h.width
	}

	h.alive = !(h.damage >= h.hp)
	return true
}

func (h *Life) Draw(screen *ebiten.Image, cam *Box) {
	ebitenutil.DrawRect(screen, h.x-1-cam.X, h.y-1-cam.Y, h.width, 2, color.RGBA{255, 0, 0, 200})
}
