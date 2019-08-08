package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

var heroIdleSprites = LoadSprites("data/hero-idle.png", 0)
var heroRunSprites = LoadSprites("data/hero-run.png", 0)
var heroJumpSprites = LoadSprites("data/hero-jump.png", 0)

type Dir int

const (
	DirRight Dir = 0
	DirLeft  Dir = 1
)

type Hero struct {
	x           float64
	y           float64
	vx          float64
	vy          float64
	inAir       bool
	jumpControl bool
	dir         Dir
	tick        int
}

func (h *Hero) Update(input Input) {

	// turn
	if input.x > 0 {
		h.dir = DirRight
	} else if input.x < 0 {
		h.dir = DirLeft
	}

	const gravity = 0.5
	const maxSpeedX = 1.5
	const maxSpeedY = 4

	// x movement
	accel := 0.5
	if h.inAir {
		accel = 0.125
	}
	if input.x != 0 {
		h.vx += float64(input.x) * accel
		h.vx = Clamp(h.vx, -maxSpeedX, maxSpeedX)
	} else if h.vx > 0 {
		h.vx = math.Max(0, h.vx-accel)
	} else if h.vx < 0 {
		h.vx = math.Min(0, h.vx+accel)
	}

	// y movement
	h.vy += gravity

	// set new pos
	h.x += h.vx
	h.y += Clamp(h.vy, -maxSpeedY, maxSpeedY)

	// TODO: collision
	if h.y > 208 {
		h.y = 208
		h.vy = 0
		h.inAir = false

		// jumping
		if input.a {
			h.inAir = true
			h.vy = -7
		}

	} else {
		h.inAir = true
	}

	h.tick++
}

func (h *Hero) Draw(screen *ebiten.Image) {
	o := ebiten.DrawImageOptions{}
	o.GeoM.Translate(-16, -24)
	if h.dir == DirLeft {
		o.GeoM.Scale(-1, 1)
	}
	o.GeoM.Translate(h.x, h.y)

	var frame *ebiten.Image

	if h.inAir {
		f := 3
		switch {
		case h.vy < -4:
			f = 0
		case h.vy < 0:
			f = 1
		case h.vy < 4:
			f = 2
		}
		frame = heroJumpSprites[f]
	} else if h.vx != 0 {
		frame = heroRunSprites[h.tick/4%8]
	} else {
		frame = heroIdleSprites[0]
	}

	screen.DrawImage(frame, &o)

	// debugging rect
	// ebitenutil.DrawRect(screen, h.x-7, h.y-19, 14, 19, color.RGBA{100, 0, 0, 100})
}
