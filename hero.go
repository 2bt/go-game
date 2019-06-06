package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var heroIdleImg *ebiten.Image
var heroRunImg *ebiten.Image

func init() {
	heroIdleImg, _, _ = ebitenutil.NewImageFromFile("data/hero-idle.png", ebiten.FilterDefault)
	heroRunImg, _, _ = ebitenutil.NewImageFromFile("data/hero-run.png", ebiten.FilterDefault)
}

type Dir int

const (
	DirRight Dir = 0
	DirLeft  Dir = 1
)

type Hero struct {
	x    float64
	y    float64
	vx   float64
	dir  Dir
	tick int
}

func (h *Hero) Update(input Input) {

	// turn
	if input.x > 0 {
		h.dir = DirRight
	} else if input.x < 0 {
		h.dir = DirLeft
	}

	var speed float64 = 1
	h.vx = float64(input.x) * speed
	h.x += h.vx
	h.y += float64(input.y) * speed

	h.tick++
}

func (h *Hero) Draw(screen *ebiten.Image) {
	o := ebiten.DrawImageOptions{}
	o.GeoM.Translate(-16, -24)
	if h.dir == DirLeft {
		o.GeoM.Scale(-1, 1)
	}
	o.GeoM.Translate(h.x, h.y)

	if h.vx != 0 {
		frameSpeed := h.tick / 4 % 8
		rect := image.Rect(frameSpeed*32, 0, (frameSpeed+1)*32, 32)
		frame := heroRunImg.SubImage(rect).(*ebiten.Image)
		screen.DrawImage(frame, &o)
	} else {
		frameSpeed := 0
		rect := image.Rect(frameSpeed*32, 0, (frameSpeed+1)*32, 32)
		frame := heroIdleImg.SubImage(rect).(*ebiten.Image)
		screen.DrawImage(frame, &o)
	}

	// debugging rect
	ebitenutil.DrawRect(screen, h.x-7, h.y-19, 14, 19, color.RGBA{100, 0, 0, 100})
}
