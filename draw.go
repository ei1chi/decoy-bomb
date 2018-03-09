package main

import (
	et "github.com/hajimehoshi/ebiten"
)

func drawAll(screen *et.Image) {

	var sp *Sprite
	op := &et.DrawImageOptions{}

	sp = sprites["star"]
	for _, f := range fragments {
		op := sp.center()
		op.GeoM.Scale(0.3, 0.3)
		op.GeoM.Rotate(f.angle)
		op.GeoM.Translate(real(f.pos), imag(f.pos))
		op.ColorM.Translate(0, 0, 0, f.speed/3.0-1.0)
		screen.DrawImage(sp.image, op)
	}

	sp = sprites["galaxy"]
	for _, g := range galaxies {
		g.draw(screen)
	}

	sp = sprites["ghost"]
	for _, g := range ghosts {
		g.draw(screen)
	}

	sp = sprites["circle"]
	for i, _ := range decoys.arr {
		d := &decoys.arr[i]
		if !d.exist {
			continue
		}
		op := sp.center()
		op.GeoM.Translate(real(d.pos), imag(d.pos))
		rate := float64(d.power) / powerMax
		alpha := -1.0 + (0.6 * rate)
		op.ColorM.Translate(0, 0, 0, alpha)
		screen.DrawImage(sp.image, op)
	}

	sp = sprites["decoy"]
	for i, _ := range decoys.arr {
		d := &decoys.arr[i]
		if !d.exist {
			continue
		}
		op := sp.center()
		op.GeoM.Translate(real(d.pos), imag(d.pos))
		rate := 1.0
		op.ColorM.Translate(0, 0, 0, rate-1.0)
		screen.DrawImage(sp.image, op)
	}

	sp = sprites["castle"]
	op = &et.DrawImageOptions{}
	op.GeoM.Translate(240-sp.halfW, 640-sp.halfH*2)
	screen.DrawImage(sp.image, op)

	sp = sprites["quit"]
	op = &et.DrawImageOptions{}
	op.GeoM.Translate(480-sp.halfW*2, 0)
	screen.DrawImage(sp.image, op)
}
