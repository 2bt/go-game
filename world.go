package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const TileSize = 16

var worldImg *ebiten.Image
var worldTileSpriteMap map[byte]*ebiten.Image

func init() {
	worldImg, _, _ = ebitenutil.NewImageFromFile("data/world.png", ebiten.FilterDefault)
	worldTileSpriteMap = map[byte]*ebiten.Image{
		'0': worldImg.SubImage(image.Rect(0, 0, TileSize, TileSize)).(*ebiten.Image),
	}

}

type World struct {
	tiles  [][]byte
	width  int
	height int
}

func NewWorld(path string) *World {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	w := &World{}
	s := bufio.NewScanner(file)
	for s.Scan() {
		w.tiles = append(w.tiles, s.Bytes())
		w.height++
		if len(s.Bytes()) > w.width {
			w.width = len(s.Bytes())
		}
	}

	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			fmt.Printf(" %02x", w.tileAt(x, y))
		}
		fmt.Println("")
	}

	return w
}

func (w *World) tileAt(x, y int) byte {
	if y < 0 || y >= len(w.tiles) || x < 0 || x >= len(w.tiles[y]) {
		return ' '
	}
	return w.tiles[y][x]
}

func (w *World) Draw(screen *ebiten.Image) {

	o := ebiten.DrawImageOptions{}
	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			t := w.tileAt(x, y)
			img, ok := worldTileSpriteMap[t]
			if !ok {
				continue
			}
			o.GeoM.Reset()
			o.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))
			screen.DrawImage(img, &o)
		}
	}

}
