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

const Gravity = 0.5
const MaxSpeedX = 1.5
const MaxSpeedY = 4

func (h *Hero) Update(input Input) {

	// turn
	if input.X > 0 {
		h.dir = DirRight
	} else if input.X < 0 {
		h.dir = DirLeft
	}

	// x movement and collision
	accel := 0.5
	if h.inAir {
		accel = 0.125
	}
	if input.X != 0 {
		h.vx += float64(input.X) * accel
		h.vx = Clamp(h.vx, -MaxSpeedX, MaxSpeedX)
	} else if h.vx > 0 {
		h.vx = math.Max(0, h.vx-accel)
	} else if h.vx < 0 {
		h.vx = math.Min(0, h.vx+accel)
	}
	h.x += h.vx

	dist := game.world.CheckCollision(AxisX, &Box{
		h.x - 7, h.y - 19, 14, 19,
	})
	if dist != 0 {
		h.x += dist
	}

	// y movement and collision
	h.vy += Gravity
	h.y += Clamp(h.vy, -MaxSpeedY, MaxSpeedY)

	dist = game.world.CheckCollision(AxisY, &Box{
		h.x - 7, h.y - 19, 14, 19,
	})
	if dist != 0 {
		h.y += dist
		h.vy = 0
		if dist < 0 {
			h.inAir = false
		}
	} else {
		h.inAir = true
	}

	if !h.inAir {
		if input.A {
			h.inAir = true
			h.vy = -7
		}
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

	// ebitenutil.DrawRect(screen, h.x-7, h.y-19, 14, 19, color.RGBA{100, 0, 0, 100})
}
