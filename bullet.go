package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Bullet struct {
	alive bool
	box   Box
	dir   Dir
	dmg   uint
	tick  int
}

func NewBullet(x, y float64, dir Dir) *Bullet {
	return &Bullet{
		alive: true,
		box:   Box{x - 4, y - 1, 8, 2},
		dir:   dir,
		dmg:   3,
	}
}

type TakeDamage interface {
	TakeDamage(dmg uint)
}

func (b *Bullet) Alive() bool { return b.alive && b.tick < 30 }

func (b *Bullet) Box() Box { return b.box }

func (b *Bullet) Update() {
	b.tick++

	// set new pos
	if b.dir == Right {
		b.box.X += 8
	} else {
		b.box.X -= 8
	}

	// collision with world
	dist := game.world.CheckCollision(AxisY, &b.box)
	if dist != 0 {
		b.alive = false
		return
	}

	// collision with mobs
	for _, m := range game.mobs {
		if !m.Alive() {
			continue
		}
		mb := m.Box()
		dist := b.box.CheckCollision(AxisY, &mb)
		if dist != 0 {
			d, ok := m.(TakeDamage)
			if ok {
				d.TakeDamage(b.dmg)
			}
			b.alive = false
			return
		}
	}
}

func (h *Bullet) Draw(screen *ebiten.Image, cam *Box) {
	ebitenutil.DrawRect(
		screen,
		h.box.X-cam.X,
		h.box.Y-cam.Y,
		h.box.W,
		h.box.H,
		color.RGBA{
			255, 255, 255, 255,
		})
}
