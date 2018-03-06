package main

import (
	"math"

	et "github.com/hajimehoshi/ebiten"
)

func drawAll(screen *et.Image) {

	var sp *Sprite

	sp = sprites["galaxy"]
	for _, g := range galaxies {
		op := sp.center()
		op.GeoM.Rotate(math.Pi * (float64(g.life) / 60))
		op.GeoM.Translate(g.pos.x, g.pos.y)
		screen.DrawImage(sp.image, op)
	}

	sp = sprites["ghost"]
	for _, g := range ghosts {
		op := sp.center()
		if g.pos.x > 200 {
			op.GeoM.Scale(-1.0, 1.0)
		}
		op.GeoM.Translate(g.pos.x, g.pos.y)
		screen.DrawImage(sp.image, op)
	}

	sp = sprites["circle"]
	for _, d := range decoys {
		op := sp.center()
		op.GeoM.Translate(d.pos.x, d.pos.y)
		rate := float64(d.power) / powerMax
		alpha := -1.0 + (0.7 * rate)
		op.ColorM.Translate(0, 0, 0, alpha)
		screen.DrawImage(sp.image, op)
	}

	sp = sprites["decoy"]
	for _, d := range decoys {
		op := sp.center()
		op.GeoM.Translate(d.pos.x, d.pos.y)
		//rate := float64(d.power) / powerMax
		rate := 1.0
		op.ColorM.Translate(0, 0, 0, rate-1.0)
		screen.DrawImage(sp.image, op)
	}
}
