package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

type Bullet struct {
	box *Box
	dir Dir
	dmg uint
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

func (b *Bullet) Box() Box { return *b.box }

func (b *Bullet) Update() bool {
	bb := b.Box()
	bb.Y += 40
	for _, m := range game.world.mobs {
		mb := m.Box()
		dist := bb.CheckCollision(AxisY, &mb) + bb.CheckCollision(AxisX, &mb)

		if dist != 0 {
			b.Hit(m)
		}
	}

	// set new pos
	if b.dir == Right {
		b.box.X += 8
	} else {
		b.box.X -= 8
	}

	dist := game.world.CheckCollision(AxisY, &Box{b.box.X - 4, b.box.Y - 1, b.box.W, b.box.H})
	if dist != 0 {
		return false
	}

	return true
}

func (h *Bullet) Draw(screen *ebiten.Image, cam *Box) {
	ebitenutil.DrawRect(
		screen,
		h.box.X-4-cam.X,
		h.box.Y-1-cam.Y,
		h.box.W,
		h.box.H,
		color.RGBA{
			255, 255, 255, 255,
		})
}

