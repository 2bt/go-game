package main

import (
	"github.com/hajimehoshi/ebiten"
	"time"
)

var mobRobotIdle = LoadSprites("data/mob-robot-idle.png", 0)

type Mob struct {
	box  *Box
	Life *Life
	dir  Dir
	tick int
}

func (m *Mob) Box() Box { return *m.box }

func (m *Mob) TakeDamage(dmg uint) {
	println("taken damage")
	m.Life.damage += dmg
}

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
	o.GeoM.Translate(-16, -24)
	o.GeoM.Translate(-cam.X, -cam.Y)
	o.GeoM.Translate(h.box.X, h.box.Y-49)

	_ = screen.DrawImage(mobRobotIdle[h.tick/4%8], &o)
}
