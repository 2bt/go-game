package main

import (
	"bufio"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

const TileSize = 16

var worldSprites = LoadSprites("data/world.png", 16)
var worldTileSpriteMap = make(map[byte]*ebiten.Image)

func init() {
	for k, v := range map[byte]int{
		'0': 0,
		'L': 2,
		'H': 4,
	} {
		worldTileSpriteMap[k] = worldSprites[v]
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
