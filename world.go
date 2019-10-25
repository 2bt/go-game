package main

import (
	"bufio"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten"
)

const TileSize = 16

var worldSprites = append(LoadSprites("data/world.png", TileSize), mobRobotIdle...)
var worldTileSpriteMap = make(map[byte]*ebiten.Image)

var collidable = map[byte]bool{
	'0': true,
	'1': true,
	'B': true,

	'^': true, // jump through platform
	'L': true, // ladder
}

func init() {
	for k, v := range map[byte]int{
		'0': 0,
		'1': 1,
		'B': 3,
		'L': 2,
		'H': 4,
		'^': 5,
		'.': 8,
	} {
		worldTileSpriteMap[k] = worldSprites[v]
	}
}

type World struct {
	tiles  [][]byte
	width  int
	height int
}

func (w *World) Load(path string, spawn func(byte, float64, float64)) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		w.tiles = append(w.tiles, s.Bytes())
		if len(s.Bytes()) > w.width {
			w.width = len(s.Bytes())
		}

		// spawn
		for x, t := range w.tiles[w.height] {
			spawn(t, float64(x*TileSize+TileSize/2), float64(w.height*TileSize+TileSize))
		}
		w.height++
	}
}

func (w *World) TileAt(x, y int) byte {
	if y < 0 || y >= len(w.tiles) || x < 0 || x >= len(w.tiles[y]) {
		return '1'
	}
	return w.tiles[y][x]
}

func (w *World) CheckCollision(axis Axis, box Box) float64 {
	return w.CheckCollisionEx(axis, box, 0)
}

func (w *World) CheckCollisionEx(axis Axis, box Box, vel float64) float64 {

	x1 := int(math.Floor(box.X / TileSize))
	x2 := int(math.Floor((box.X + box.W) / TileSize))
	y1 := int(math.Floor(box.Y / TileSize))
	y2 := int(math.Floor((box.Y + box.H) / TileSize))

	var dist float64
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			t := w.TileAt(x, y)

			// ignore background tile
			if !collidable[t] {
				continue
			}
			if (t == '^' ||
				t == 'L') &&
				(axis != AxisY || vel < 0) {
				continue
			}

			// only top end of ladder has collision
			if t == 'L' && w.TileAt(x, y-1) == 'L' {
				continue
			}

			// check collision with tile box
			d := box.CheckCollision(axis, Box{
				float64(x * TileSize),
				float64(y * TileSize),
				TileSize,
				TileSize,
			})

			if t == '^' || t == 'L' {
				if d >= 0 || vel < -d {
					continue
				}
			}

			if math.Abs(d) > math.Abs(dist) {
				dist = d
			}
		}
	}
	return dist
}

func (w *World) Draw(screen *ebiten.Image, cam Box) {

	x1 := int(math.Floor(cam.X / TileSize))
	x2 := int(math.Floor((cam.X + cam.W) / TileSize))
	y1 := int(math.Floor(cam.Y / TileSize))
	y2 := int(math.Floor((cam.Y + cam.H) / TileSize))

	o := ebiten.DrawImageOptions{}
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			t := w.TileAt(x, y)

			img, ok := worldTileSpriteMap[t]
			if !ok {
				continue
			}
			o.GeoM.Reset()
			o.GeoM.Translate(float64(x*TileSize)-cam.X, float64(y*TileSize)-cam.Y)
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
