package main

import (
	"math"
	"math/rand"

	et "github.com/hajimehoshi/ebiten"
)

var (
	galaxies []Galaxy
)

func processGalaxies() {
	for _, g := range galaxies {
		g.update()
	}
}

// Base Class
//=========================

type Galaxy interface {
	update()
	draw(screen *et.Image)
	isDead() bool
}

type GalaxyBase struct {
	count int
	dead  bool
	pos   complex128
}

func (g *GalaxyBase) isDead() bool {
	return g.dead
}

func (g *GalaxyBase) drawSimple(sc *et.Image) {
	sp := sprites["galaxy"]
	op := sp.center()
	op.GeoM.Rotate(math.Pi * (float64(g.count) / 60))
	op.GeoM.Translate(real(g.pos), imag(g.pos))
	sc.DrawImage(sp.image, op)
}

func (g *GalaxyBase) checkArea() {
	x := real(g.pos)
	if x < -50 || 530 < x {
		g.dead = true
	}
	y := imag(g.pos)
	if y < -50 || 690 < y {
		g.dead = true
	}
}

// Sub Classes
//=========================

type NormalGx struct {
	GalaxyBase
}

func (g *NormalGx) draw(sc *et.Image) {
	g.drawSimple(sc)
}

func (g *NormalGx) update() {
	switch {
	case g.count == 0: // init
	case g.count < 180:
		if g.count%80 == 0 {
			dir := powi(around * rand.Float64()) // any direction
			const ways = 8
			c := powi(around / ways)
			for n := 0; n < ways; n++ {
				ghost := &NormalGs{}
				ghost.pos = g.pos
				ghost.vec = dir
				ghosts = append(ghosts, ghost)
				dir *= c
			}
		}
	default:
		g.dead = true
	}
	g.checkArea()
	g.count += 1
}
