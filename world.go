package main

import (
	"bufio"
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

	worldTileSpriteMap = make(map[byte]*ebiten.Image)

	for k, v := range map[byte]int{
		'0': 0x0000,
		'L': 0x0002,
		'H': 0x0004,
	} {
		x := v & 0xff
		y := (v >> 8) & 0xff
		rect := image.Rect(x*TileSize, y*TileSize, (x+1)*TileSize, (y+1)*TileSize)
		worldTileSpriteMap[k] = worldImg.SubImage(rect).(*ebiten.Image)
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
