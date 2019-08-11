package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

var heroIdleSprites = LoadSprites("data/hero-idle.png", 0)
var heroRunSprites = LoadSprites("data/hero-run.png", 0)
var heroJumpSprites = LoadSprites("data/hero-jump.png", 0)
var heroClimbSprites = LoadSprites("data/hero-climb.png", 0)

type Hero struct {
	x           float64
	y           float64
	vx          float64
	vy          float64
	dir         Dir
	climbing    bool
	inAir       bool
	jumpControl bool
	oldJump     bool
	shootDelay  int
	tick        int
	Life        *Life
}

const Gravity = 0.5
const MaxSpeedX = 1.5
const MaxSpeedY = 3

func (h *Hero) Update(input Input) {

	// turn
	if input.x > 0 {
		h.dir = Right
	} else if input.x < 0 {
		h.dir = Left
	}

	if h.climbing {

		m := math.Mod(h.x, TileSize)
		h.x += Clamp(8-m, -0.5, 0.5)

		h.vy = float64(input.y)
		h.y += h.vy

		dist := game.world.CheckCollision(AxisY, &Box{
			h.x - 7, h.y - 19, 14, 19,
		})
		if dist != 0 {
			h.y += dist
			h.vy = 0
			if dist < 0 {
				h.inAir = false
				h.climbing = false
			}
		}

		x := int(h.x / TileSize)
		y := int((h.y - 0.5) / TileSize)
		t := game.world.TileAt(x, y)
		if t != 'L' {
			if h.dir == Right {
				h.x += 1
			} else {
				h.x -= 1
			}
			h.climbing = false
		}

		if input.jump {
			h.inAir = true
			h.climbing = false
		}

	} else {

		// x movement and collision
		accel := 0.5
		if h.inAir {
			accel = 0.25
		}
		if input.x != 0 {
			h.vx += float64(input.x) * accel
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

		if h.inAir {
			// jump higher
			if h.jumpControl {
				if !input.jump && h.vy < -1 {
					h.vy = -1
					h.jumpControl = false
				}
				if !input.jump || h.vy > -1 {
					h.jumpControl = false
				}
			}

		} else {
			// jump
			if input.jump && !h.oldJump {
				h.inAir = true
				h.jumpControl = true
				h.vy = -9
			}
		}

		// check for latter
		if input.y == -1 {
			x := int(h.x / TileSize)
			y := int((h.y - 0.5) / TileSize)
			t := game.world.TileAt(x, y)
			m := math.Mod(h.x, TileSize)
			if t == 'L' && m > 5 && m < 11 {
				h.inAir = false
				h.climbing = true
				h.vx = 0
				h.vy = 0
			}
		}

	}

	h.oldJump = input.jump

	// fire a bullet
	if input.shoot && h.shootDelay <= 0 {
		game.AddBullet(&Bullet{&Box{h.x, h.y - 11, 8, 2}, h.dir, 3})
		h.shootDelay = 10
	}
	if h.shootDelay > 0 {
		h.shootDelay--
	}

	h.Life.Update(h.x-3, h.y-2)
	h.tick++
}

func (h *Hero) Draw(screen *ebiten.Image, cam *Box) {
	h.Life.Draw(screen, cam)
	o := ebiten.DrawImageOptions{}
	o.GeoM.Translate(-16, -24)
	if h.dir == Left {
		o.GeoM.Scale(-1, 1)
	}
	o.GeoM.Translate(-cam.X, -cam.Y)
	o.GeoM.Translate(h.x, h.y)

	var frame *ebiten.Image

	if h.climbing {
		frame = heroClimbSprites[0]
		if h.vy != 0 {
			frame = heroClimbSprites[h.tick/4%4]
		}
	} else if h.inAir {
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