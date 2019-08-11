package main

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

var mobRobotIdle = LoadSprites("data/mob-robot-idle.png", 0)
var mobRobotDie = LoadSprites("data/mob-robot-die1.png", 0)

type Mob struct {
	box  Box
	Life *Life
	dir  Dir
	tick int
}

func NewMob(x, y float64) *Mob {
	return &Mob{
		box: Box{
			X: x - 8,
			Y: y - 24,
			W: 16,
			H: 24,
		},
		Life: &Life{hp: 20, x: x + 100, y: y, ownerSize: 70, alive: true, xOffset: 10},
	}
}

func (m *Mob) ToRemove() bool      { return !m.Life.alive }
func (m *Mob) Box() Box            { return m.box }
func (m *Mob) TakeDamage(dmg uint) { m.Life.damage += dmg }

func (m *Mob) Update() bool {
	if time.Now().Unix()%2 == 0 {
		m.dir = Left
		m.box.X -= 1
	} else {
		m.dir = Right
		m.box.X += 1
	}

	m.Life.Update(m.box.X+8, m.box.Y-6)
	m.tick++
	return true
}

func (h *Mob) Draw(screen *ebiten.Image, cam *Box) {
	h.Life.Draw(screen, cam)
	o := ebiten.DrawImageOptions{}
	o.GeoM.Translate(-32+8, -44+24)
	o.GeoM.Translate(-cam.X, -cam.Y)
	o.GeoM.Translate(h.box.X, h.box.Y)

	if h.Life.alive {
		_ = screen.DrawImage(mobRobotIdle[h.tick/4%8], &o)
		return
	}
	//todo
	// how to change these frames slower
	// at the moment cannot see the whole animation
	for _, frame := range mobRobotDie {
		_ = screen.DrawImage(frame, &o)
	}

	// // rect for debugging
	// ebitenutil.DrawRect(
	// 	screen,
	// 	h.box.X-cam.X,
	// 	h.box.Y-cam.Y,
	// 	h.box.W,
	// 	h.box.H,
	// 	color.RGBA{
	// 		0, 255, 0, 100,
	// 	})
}
