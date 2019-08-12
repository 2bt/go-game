package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

var mobRobotIdle = LoadSprites("data/mob-robot-idle.png", 0)
var mobRobotDie = LoadSprites("data/mob-robot-die1.png", 0)

type Mob struct {
	box         Box
	Life        *Life
	vx          float64
	walkCounter int
	tick        int
	deathAnim   bool
	deathTick   int
}

func NewMob(x, y float64) *Mob {
	return &Mob{
		box: Box{
			X:          x - 8,
			Y:          y - 24,
			W:          16,
			H:          24,
			collidable: true,
		},
		Life:        &Life{hp: 20, x: x + 100, y: y, ownerSize: 70, alive: true, xOffset: 10},
		walkCounter: rand.Intn(90),
		vx:          1,
	}
}

func (m *Mob) ToRemove() bool      { return !m.Life.alive && !m.deathAnim }
func (m *Mob) Box() Box            { return m.box }
func (m *Mob) TakeDamage(dmg uint) { m.Life.damage += dmg }

func (m *Mob) Update() bool {
	if m.walkCounter > 0 {
		m.walkCounter--
	} else {
		m.walkCounter = 90 + rand.Intn(30)
		m.vx *= -1
	}
	m.box.X += m.vx

	m.Life.Update(m.box.X+8, m.box.Y-6)
	m.tick++

	if !m.deathAnim {
		m.deathAnim = !m.Life.alive
	}
	if m.deathAnim {
		m.box.collidable = false
		m.deathTick++
	}
	return true
}

//todo lore idea
// save the princess/prisoner but she turns into a prophet
//  use the socrates bridge paradox
// https://en.wikipedia.org/wiki/Buridan%27s_bridge
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

	if h.deathTick >= 176 {
		h.deathTick = 0
		h.deathAnim = false
		return
	}

	_ = screen.DrawImage(mobRobotDie[h.deathTick], &o)

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
