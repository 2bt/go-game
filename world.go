package main

import (
	"bufio"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten"
)

const TileSize = 16

var worldSprites = LoadSprites("data/world.png", TileSize)
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

func (w *World) TileAt(x, y int) byte {
	if y < 0 || y >= len(w.tiles) || x < 0 || x >= len(w.tiles[y]) {
		return ' '
	}
	return w.tiles[y][x]
}

func (w *World) CheckCollision(axis Axis, box *Box) float64 {

	x1 := int(math.Floor(box.X / TileSize))
	x2 := int(math.Floor((box.X + box.W) / TileSize))
	y1 := int(math.Floor(box.Y / TileSize))
	y2 := int(math.Floor((box.Y + box.H) / TileSize))

	var dist float64
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			t := w.TileAt(x, y)

			// ignore background tile
			if t == ' ' || t == 'L' {
				continue
			}

			// check collision with tile box
			d := box.CheckCollision(axis, &Box{
				float64(x * TileSize),
				float64(y * TileSize),
				TileSize,
				TileSize,
			})
			if math.Abs(d) > math.Abs(dist) {
				dist = d
			}

		}
	}
	return dist
}

func (w *World) Draw(screen *ebiten.Image) {
	o := ebiten.DrawImageOptions{}
	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			t := w.TileAt(x, y)
			img, ok := worldTileSpriteMap[t]
			if !ok {
				continue
			}
			o.GeoM.Reset()
			o.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))
			screen.DrawImage(img, &o)
		}
	}

	/*
		// debug
		h := game.hero
		box := Box{
			h.x - 7, h.y - 19, 14, 19,
		}
		x1 := int(math.Floor(box.X / TileSize))
		x2 := int(math.Floor((box.X + box.W) / TileSize))
		y1 := int(math.Floor(box.Y / TileSize))
		y2 := int(math.Floor((box.Y + box.H) / TileSize))

		for y := y1; y <= y2; y++ {
			for x := x1; x <= x2; x++ {
				t := w.TileAt(x, y)
				if t == ' ' {
					continue
				}
				b := Box{
					float64(x * TileSize),
					float64(y * TileSize),
					TileSize,
					TileSize,
				}
				ebitenutil.DrawRect(screen, b.X, b.Y, b.W, b.H, color.RGBA{100, 100, 0, 100})

			}
		}
	*/
}
