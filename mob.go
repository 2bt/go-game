package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

var (
	mobRobotIdle = LoadSprites("data/mob-robot-idle.png", 0)
	mobRobotDie  = LoadSprites("data/mob-robot-die1.png", 0)

	//make sure that Mob implements the TakeDamage interface
	_ TakeDamage = (*Mob)(nil)
)

type Mob struct {
	box         Box
	life        *Life
	vx          float64
	walkCounter int
	tick        int
}

func NewMob(x, y float64) *Mob {
	return &Mob{
		box: Box{
			X: x - 8,
			Y: y - 24,
			W: 16,
			H: 24,
		},
		life:        &Life{hp: 20, x: x + 100, y: y, ownerSize: 70, alive: true, xOffset: 10},
		walkCounter: rand.Intn(90),
		vx:          1,
	}
}

func (m *Mob) Alive() bool         { return m.life.alive }
func (m *Mob) Box() Box            { return m.box }
func (m *Mob) TakeDamage(dmg uint) { m.life.damage += dmg }

func (m *Mob) Update() {
	if m.walkCounter > 0 {
		m.walkCounter--
	} else {
		m.walkCounter = 90 + rand.Intn(30)
		m.vx *= -1
	}
	m.box.X += m.vx

	m.life.Update(m.box.X+8, m.box.Y-6)
	m.tick++

	if !m.life.alive {
		// spawn particle
		game.particles = append(game.particles, NewMobDeathParticle(m.box.X, m.box.Y))
	}
}

//todo lore idea
// save the princess/prisoner but she turns into a prophet
//  use the socrates bridge paradox
// https://en.wikipedia.org/wiki/Buridan%27s_bridge
func (h *Mob) Draw(screen *ebiten.Image, cam *Box) {
	h.life.Draw(screen, cam)
	o := ebiten.DrawImageOptions{}
	o.GeoM.Translate(-32+8, -44+24)
	o.GeoM.Translate(-cam.X, -cam.Y)
	o.GeoM.Translate(h.box.X, h.box.Y)

	_ = screen.DrawImage(mobRobotIdle[h.tick/4%8], &o)
}

type MobDeathParticle struct {
	x, y float64
	tick int
}

func NewMobDeathParticle(x, y float64) *MobDeathParticle {
	return &MobDeathParticle{
		x: x,
		y: y,
	}
}

func (p *MobDeathParticle) Update() {
	p.tick++

}
func (p *MobDeathParticle) Draw(screen *ebiten.Image, cam *Box) {
	o := ebiten.DrawImageOptions{}
	o.GeoM.Translate(-32+8, -44+24)
	o.GeoM.Translate(-cam.X, -cam.Y)
	o.GeoM.Translate(p.x, p.y)
	_ = screen.DrawImage(mobRobotDie[p.tick/2], &o)
}
func (p *MobDeathParticle) Box() Box { return Box{} }
func (p *MobDeathParticle) Alive() bool {
	return p.tick < 44

}
