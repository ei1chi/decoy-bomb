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
		g.pos = complex(rand.Float64()*480, rand.Float64()*640)
		galaxies = append(galaxies, g)
	}
}
