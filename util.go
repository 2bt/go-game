package main

import (
	"image"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Box struct {
	X float64
	Y float64
	W float64
	H float64
}

type Axis int

const (
	AxisX Axis = 0
	AxisY Axis = 1
)

func Clamp(x, a, b float64) float64 {
	return math.Max(a, math.Min(b, x))
}

func LoadSprites(path string, size int) []*ebiten.Image {
	var img, _, err = ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	w, h := img.Size()
	if size == 0 {
		size = h
	}
	var sprites []*ebiten.Image

	for y := 0; y < h/size; y++ {
		for x := 0; x < w/size; x++ {
			rect := image.Rect(x*size, y*size, x*size+size, y*size+size)
			frame := img.SubImage(rect).(*ebiten.Image)
			sprites = append(sprites, frame)
		}
	}
	return sprites
}

func (a *Box) CheckCollision(axis Axis, b *Box) float64 {
	if a.X >= b.X+b.W || a.Y >= b.Y+b.H || a.X+a.W <= b.X || a.Y+a.H <= b.Y {
		return 0
	}
	var v, w float64
	if axis == AxisX {
		v = b.X + b.W - a.X
		w = b.X - a.X - a.W
	} else {
		v = b.Y + b.H - a.Y
		w = b.Y - a.Y - a.H
	}
	if math.Abs(v) < math.Abs(w) {
		return v
	}
	return w
}
