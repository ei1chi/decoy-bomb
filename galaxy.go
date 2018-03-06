package main

import "math/rand"

type Galaxy struct {
	life int // from 0
	pos  Point
}

var (
	galaxies []*Galaxy
	genCount = 100
)

func processGalaxies() {

	for _, g := range galaxies {
		g.life += 1
	}

	genCount -= 1
	if genCount < 0 {
		genCount = 100
		x, y := rand.Intn(480), rand.Intn(640)
		galaxies = append(galaxies, &Galaxy{pos: Point{float64(x), float64(y)}})
	}
}
