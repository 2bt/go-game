package main

import (
	"image"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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
