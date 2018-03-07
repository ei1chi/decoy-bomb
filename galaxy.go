package main

import (
	"math"

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
	pos   Point
}

func (g *GalaxyBase) isDead() bool {
	return g.dead
}

func (g *GalaxyBase) drawSimple(sc *et.Image) {
	sp := sprites["galaxy"]
	op := sp.center()
	op.GeoM.Rotate(math.Pi * (float64(g.count) / 60))
	op.GeoM.Translate(g.pos.x, g.pos.y)
	sc.DrawImage(sp.image, op)
}

func (g *GalaxyBase) checkArea() {
	if g.pos.x < -50 || 530 < g.pos.x {
		g.dead = true
	}
	if g.pos.y < -50 || 690 < g.pos.y {
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
	case g.count%80 == 0:
		ghost := &NormalGs{}
		ghost.pos = g.pos
		ghosts = append(ghosts, ghost)
	case g.count > 180:
		g.dead = true
	}
	g.checkArea()
	g.count += 1
}
