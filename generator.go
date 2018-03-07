package main

import "math/rand"

var (
	interval = 150
	genCount = interval
)

func processGenerator() {
	genCount -= 1
	if genCount < 0 {
		genCount = interval
		g := &NormalGx{}
		x, y := rand.Intn(480), rand.Intn(640)
		g.pos = Point{float64(x), float64(y)}
		galaxies = append(galaxies, g)
	}
}
