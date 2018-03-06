package main

import "math/rand"

type Galaxy struct {
	update        func(*Galaxy)
	life, lifeMax int
	pos           Point
}

func normalGalaxy(g *Galaxy) {
	if g.life%90 == 0 {
		v := Point{1, 0}
		ghosts = append(ghosts, &Ghost{update: normalGhost, pos: g.pos, vec: v})
	}
}

var (
	galaxies []*Galaxy
	genCount = 100
)

func processGalaxies() {

	for _, g := range galaxies {
		g.life += 1
		g.update(g)
	}

	genCount -= 1
	if genCount < 0 {
		genCount = 100
		g := &Galaxy{}
		x, y := rand.Intn(480), rand.Intn(640)
		g.pos = Point{float64(x), float64(y)}
		g.lifeMax = 120
		g.update = normalGalaxy
		galaxies = append(galaxies, g)
	}
}
