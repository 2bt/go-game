package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Bullet struct {
	box Box
	dir Dir
	dmg uint
}

func NewBullet(x, y float64, dir Dir) *Bullet {
	return &Bullet{
		Box{x - 4, y - 1, 8, 2},
		dir, 3,
	}
}

type TakeDamage interface {
	TakeDamage(dmg uint)
}

func (b *Bullet) Hit(e Entity) {
	d, ok := e.(TakeDamage)
	if ok {
		d.TakeDamage(b.dmg)
	}
}

func (b *Bullet) Box() Box { return b.box }

func (b *Bullet) Update() bool {

	// set new pos
	if b.dir == Right {
		b.box.X += 8
	} else {
		b.box.X -= 8
	}

	// collision with world
	dist := game.world.CheckCollision(AxisY, &b.box)
	if dist != 0 {
		return false
	}

	// collision with mobs
	for _, m := range game.world.mobs {
		mb := m.Box()
		dist := b.box.CheckCollision(AxisY, &mb)
		if dist != 0 {
			b.Hit(m)
		}
	}

	return true
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
