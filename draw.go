package main

import (
	et "github.com/hajimehoshi/ebiten"
)

func drawAll(screen *et.Image) {

	var sp *Sprite
	op := &et.DrawImageOptions{}

	sp = sprites["galaxy"]
	for _, g := range galaxies {
		g.draw(screen)
	}

	sp = sprites["ghost"]
	for _, g := range ghosts {
		g.draw(screen)
	}

	sp = sprites["circle"]
	for _, d := range decoys {
		op := sp.center()
		op.GeoM.Translate(d.pos.x, d.pos.y)
		rate := float64(d.power) / powerMax
		alpha := -1.0 + (0.6 * rate)
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

	sp = sprites["castle"]
	op = &et.DrawImageOptions{}
	op.GeoM.Translate(240-sp.halfW, 640-sp.halfH*2)
	screen.DrawImage(sp.image, op)
}
