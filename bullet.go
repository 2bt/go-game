package main

import (
	"image/color"
	"math/rand"

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
		for i := 0; i < 20; i++ {
			game.AddParticle(NewSparkParticle(b.box.X+8+dist, b.box.Y+1))
		}
		b.alive = false
		return
	}

	// collision with mobs
	for _, m := range game.mobs {
		if !m.Alive() {
			continue
		}
		mb := m.Box()
		dist := b.box.CheckCollision(AxisX, &mb)
		if dist != 0 {
			d, ok := m.(TakeDamage)
			if ok {
				d.TakeDamage(b.dmg)
			}
			for i := 0; i < 20; i++ {
				game.AddParticle(NewSparkParticle(b.box.X+8+dist, b.box.Y+1))
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

type SparkParticle struct {
	box  Box
	vx   float64
	vy   float64
	tick int
}

func NewSparkParticle(x, y float64) *SparkParticle {
	return &SparkParticle{
		box:  Box{x - 1, y - 1, 2, 2},
		vx:   (rand.Float64() - 0.5) * 5,
		vy:   (rand.Float64() - 0.5) * 5,
		tick: rand.Intn(20) + 5,
	}
}

func (p *SparkParticle) Update() {
	p.tick--
	p.box.X += p.vx
	dist := game.world.CheckCollision(AxisX, &p.box)
	if dist != 0 {
		p.box.X += dist
		p.vx *= 1
	}
	p.vy += Gravity
	p.box.Y += Clamp(p.vy, -MaxSpeedY, MaxSpeedY)
	dist = game.world.CheckCollision(AxisY, &p.box)
	if dist != 0 {
		p.box.Y += dist
		p.vy *= 1
	}

}
func (p *SparkParticle) Draw(screen *ebiten.Image, cam *Box) {
	ebitenutil.DrawRect(
		screen,
		p.box.X-cam.X,
		p.box.Y-cam.Y,
		p.box.W,
		p.box.H,
		color.RGBA{
			255, 255, 255, 100,
		})
}
func (p *SparkParticle) Box() Box { return Box{} }
func (p *SparkParticle) Alive() bool {
	return p.tick > 0

}
