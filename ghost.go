package main

import (
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
	hit(complex128, StarStates)
}

type GhostBase struct {
	count               int
	dead                bool
	pos, vec, targetPos complex128
	targetState         StarStates
}

func (g *GhostBase) isDead() bool {
	return g.dead
}

func (g *GhostBase) collInfo() (pos complex128, r float64) {
	return g.pos, 16
}

func (g *GhostBase) hit(spos complex128, state StarStates) {
	if state == starBlasting {
		g.dead = true
	}
	g.targetPos = spos
	g.targetState = state
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
		g.speed = 1
	//case g.count < 60:
	//	g.speed = 2.0 - (float64(g.count)/60)*1.0
	//case g.count < 90:
	//	g.speed = 1.0 + float64(g.count-60)/30
	case g.count < 600:
		g.speed = 1
	default:
		g.dead = true
	}

	switch g.targetState {
	case starMoving:
		diff := g.targetPos - g.pos
		diff /= complex(cmplx.Abs(diff), 0)
		g.vec -= diff / 9
	}
	g.vec *= complex(g.speed/cmplx.Abs(g.vec), 0)
	g.pos += g.vec
	g.checkArea()
	g.count += 1
}
