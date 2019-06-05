package main

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var heroIdleImg *ebiten.Image
var heroRunImg *ebiten.Image

func init() {
	heroIdleImg, _, _ = ebitenutil.NewImageFromFile("data/hero-idle2.png", ebiten.FilterDefault)
	heroRunImg, _, _ = ebitenutil.NewImageFromFile("data/hero-run.png", ebiten.FilterDefault)
}

type Hero struct {
	x    float64
	y    float64
	vx   float64
	dir  float64
	tick int
}

func (h *Hero) Update(input Input) {

	var speed float64 = 1

	if input.x != 0 {
		h.dir = float64(input.x)
	}

	h.vx = float64(input.x) * speed
	h.x += h.vx
	h.y += float64(input.y) * speed

	h.tick++
}

func (h *Hero) Draw(screen *ebiten.Image) {
	o := ebiten.DrawImageOptions{}
	if h.dir == -1 {
		o.GeoM.Scale(h.dir, 1)
		o.GeoM.Translate(32, 0)
	}
	o.GeoM.Translate(h.x, h.y)

	if h.vx != 0 {
		f := h.tick / 4 % 8
		rect := image.Rect(f*32, 0, (f+1)*32, 32)
		frame := heroRunImg.SubImage(rect).(*ebiten.Image)
		screen.DrawImage(frame, &o)
	} else {
		f := 0
		rect := image.Rect(f*32, 0, (f+1)*32, 32)
		frame := heroIdleImg.SubImage(rect).(*ebiten.Image)
		screen.DrawImage(frame, &o)
	}

}
