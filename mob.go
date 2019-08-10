package main

import (
	"github.com/hajimehoshi/ebiten"
	"time"
)

var mobRobotIdle = LoadSprites("data/mob-robot-idle.png", 0)

type Mob struct {
	x    float64
	y    float64
	dir  Dir
	tick int
	Life *Life
}

func (m *Mob) Update() bool {
	if time.Now().Unix()%2 == 0 {
		m.dir = Left
		m.x -= 1
	} else {
		m.dir = Right
		m.x += 1
	}

	m.Life.x, m.Life.y = m.x, m.y
	m.tick++
	return true
}

func (h *Mob) Draw(screen *ebiten.Image, cam *Box) {
	h.Life.Draw(screen, cam)
	o := ebiten.DrawImageOptions{}
	o.GeoM.Translate(-16, -24)
	o.GeoM.Translate(-cam.X, -cam.Y)
	o.GeoM.Translate(h.x, h.y-49)
	_ = screen.DrawImage(mobRobotIdle[h.tick/4%8], &o)
}
