package main

import "github.com/hajimehoshi/ebiten"

type Entity interface {
	Update() bool
	Draw(screen *ebiten.Image, cam *Box)
	Box() Box
}

type Entities []Entity

func (es *Entities) Update() {
	i := 0
	for _, e := range *es {
		if e.Update() {
			(*es)[i] = e
			i++
		}
	}
	*es = (*es)[:i]
}

func (es Entities) Draw(screen *ebiten.Image, cam *Box) {
	for _, e := range es {
		e.Draw(screen, cam)
	}
}
