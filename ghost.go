package main

import (
	"log"
	"math/cmplx"

	et "github.com/hajimehoshi/ebiten"
)

var (
	ghosts []Ghost
)

func processGhosts() {
	for _, g := range ghosts {
		g.update()
	}
}

// Base Class
//=========================

type Ghost interface {
	update()
	draw(screen *et.Image)
	isDead() bool
	collInfo() (pos complex128, r float64)
	hit(int)
}

type GhostBase struct {
	count    int
	dead     bool
	pos, vec complex128
	leadTo   int // player base(-1) or decoy(0~)
}

func (g *GhostBase) isDead() bool {
	return g.dead
}

func (g *GhostBase) collInfo() (pos complex128, r float64) {
	return g.pos, 32
}

func (g *GhostBase) hit(id int) {
	if g.leadTo != id {
		log.Print("hit")
	}
	g.leadTo = id
}

func (g *GhostBase) drawSimple(sc *et.Image) {
	sp := sprites["ghost"]
	op := sp.center()
	if real(g.vec) < 0 {
		op.GeoM.Scale(-1.0, 1.0)
	}
	op.GeoM.Translate(real(g.pos), imag(g.pos))
	sc.DrawImage(sp.image, op)
}

func (g *GhostBase) checkArea() {
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

type NormalGs struct {
	GhostBase
	speed float64
}

func (g *NormalGs) draw(sc *et.Image) {
	g.drawSimple(sc)
}

func (g *NormalGs) update() {
	switch {
	case g.count == 0: // init
	case g.count < 60:
		g.speed = 2.0 - (float64(g.count)/60)*1.0
	case g.count < 90:
		g.speed = 1.0 + float64(g.count-60)/30
	case g.count < 180:
		g.speed = 2
	default:
		g.dead = true
	}

	to := decoys.arr[g.leadTo]
	if to.exist {
		diff := to.pos - g.pos
		diff /= complex(cmplx.Abs(diff), 0)
		g.vec += diff * 0.7
		g.vec = g.vec / complex(cmplx.Abs(g.vec)*g.speed, 0)
	}
	g.pos += g.vec
	g.checkArea()
	g.count += 1
}
