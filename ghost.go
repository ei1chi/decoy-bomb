package main

import (
	"log"

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
	collInfo() (x, y, r float64)
	hit(int)
}

type GhostBase struct {
	count        int
	dead         bool
	pos          Point
	speed, angle float64
	leadTo       int // player base(-1) or decoy(0~)
}

func (g *GhostBase) isDead() bool {
	return g.dead
}

func (g *GhostBase) collInfo() (x, y, r float64) {
	return g.pos.x, g.pos.y, 32
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
	if g.pos.x > 200 {
		op.GeoM.Scale(-1.0, 1.0)
	}
	op.GeoM.Translate(g.pos.x, g.pos.y)
	sc.DrawImage(sp.image, op)
}

func (g *GhostBase) checkArea() {
	if g.pos.x < -50 || 530 < g.pos.x {
		g.dead = true
	}
	if g.pos.y < -50 || 690 < g.pos.y {
		g.dead = true
	}
}

// Sub Classes
//=========================

type NormalGs struct {
	GhostBase
}

func (g *NormalGs) draw(sc *et.Image) {
	g.drawSimple(sc)
}

func (g *NormalGs) update() {
	switch {
	case g.count == 0: // init
	case g.count < 30:
		g.pos.x += 1
	case 120 < g.count:
		g.dead = true
	}
	//g.pos.x += g.vec.x
	//g.pos.y += g.vec.y
	g.checkArea()
	g.count += 1
}
