package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

var heroIdleSprites = LoadSprites("data/hero-idle.png", 0)
var heroRunSprites = LoadSprites("data/hero-run.png", 0)
var heroJumpSprites = LoadSprites("data/hero-jump.png", 0)
var heroClimbSprites = LoadSprites("data/hero-climb.png", 0)

type HeroState int

const (
	OnGround HeroState = iota
	InAir
	Climbing
)

type Hero struct {
	x           float64
	y           float64
	vx          float64
	vy          float64
	dir         Dir
	state       HeroState
	jumpControl bool
	oldJump     bool
	shootDelay  int
	tick        int
	life        *Life
}

const Gravity = 0.5
const MaxSpeedX = 1.5
const MaxSpeedY = 3

func NewHero(x, y float64) *Hero {
	return &Hero{
		x:    x,
		y:    y,
		life: &Life{hp: 100, x: x, y: y, ownerSize: 35, xOffset: 1, yOffset: 20},
	}
}

func (h *Hero) Box() Box {
	return Box{h.x - 7, h.y - 19, 14, 19}
}

func (h *Hero) Update(input Input) {

	// turn
	if input.x > 0 {
		h.dir = Right
	} else if input.x < 0 {
		h.dir = Left
	}

	if h.state == Climbing {

		// snap to ladder
		m := math.Mod(h.x, TileSize)
		h.x += Clamp(8-m, -1, 1)

		// move up/down
		h.vy = float64(input.y)
		h.y += h.vy

		// collision
		dist := game.world.CheckCollision(AxisY, h.Box())
		if dist != 0 {
			h.y += dist
			h.vy = 0
			if dist < 0 {
				h.state = OnGround
			}
		}

		// check if we're still on a ladder
		x := int(h.x / TileSize)
		y := int((h.y - 0.5) / TileSize)
		t := game.world.TileAt(x, y)
		if t != 'L' {
			if h.vy < 0 {
				// we reached the top end
				dist := game.world.CheckCollisionEx(AxisY, h.Box(), -h.vy)
				h.y += dist
				h.vy = 0
				h.state = OnGround
			} else {
				h.state = InAir
			}
		}

		// let go of ladder
		if input.jump && !h.oldJump {
			h.state = InAir
		}

	} else {

		// NOTE: movement and collision are handled the same InAir and OnGround

		// x movement and collision
		accel := 0.5
		if h.state == InAir {
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

		dist := game.world.CheckCollision(AxisX, h.Box())
		if dist != 0 {
			h.x += dist
		}

		// y movement and collision
		h.vy += Gravity
		var vy = Clamp(h.vy, -MaxSpeedY, MaxSpeedY)
		h.y += vy

		dist = game.world.CheckCollisionEx(AxisY, h.Box(), vy)
		if dist != 0 {
			h.y += dist
			h.vy = 0
			if dist < 0 {
				h.state = OnGround
			}
		} else {
			h.state = InAir
		}

		if h.state == InAir {
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

		} else { // OnGround

			// jump
			if input.jump && !h.oldJump {
				h.state = InAir
				h.jumpControl = true
				h.vy = -9
			}
		}

		// check for latter
		if input.y != 0 {
			y := h.y + float64(input.y)
			ix := int(h.x / TileSize)
			iy := int(y / TileSize)
			t := game.world.TileAt(ix, iy)
			m := math.Mod(h.x, TileSize)
			if t == 'L' && m > 5 && m < 11 {
				h.state = Climbing
				h.vx = 0
				h.vy = 0
			}
		}

	}

	h.oldJump = input.jump

	// fire a bullet
	if input.shoot && h.shootDelay <= 0 {
		game.AddBullet(NewBullet(h.x, h.y-11, h.dir))
		h.shootDelay = 10
	}
	if h.shootDelay > 0 {
		h.shootDelay--
	}

	h.life.Update(h.x-3, h.y-2)
	h.tick++
}

func (h *Hero) Draw(screen *ebiten.Image, cam Box) {
	h.life.Draw(screen, cam)
	o := ebiten.DrawImageOptions{}
	o.GeoM.Translate(-16, -24)
	if h.dir == Left {
		o.GeoM.Scale(-1, 1)
	}
	o.GeoM.Translate(-cam.X, -cam.Y)
	o.GeoM.Translate(h.x, h.y)

	var frame *ebiten.Image

	if h.state == Climbing {

		frame = heroClimbSprites[0]
		if h.vy != 0 {
			frame = heroClimbSprites[h.tick/4%4]
		}
	} else if h.state == InAir {
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
