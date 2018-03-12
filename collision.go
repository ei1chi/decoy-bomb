package main

import (
	"math"
)

func collision() {
	for i, _ := range stars.arr {
		s := &stars.arr[i]
		if !s.exist {
			continue
		}
		for _, g := range ghosts {
			hit := collStarGhost(s, g)
			if hit {
				g.hit(s.pos, s.state)
			} else {
				g.hit(0, noStar)
			}
		}
	}
}

func collStarGhost(s *Star, g Ghost) bool {
	var sr float64
	switch s.state {
	case starMoving:
		sr = 64
	case starFired:
		return false // 強制終了
	case starBlasting:
		sr = 120
	}

	gpos, gr := g.collInfo()
	diff := s.pos - gpos
	distsq := math.Pow(real(diff), 2) + math.Pow(imag(diff), 2)
	limitsq := math.Pow(sr+gr, 2)
	return limitsq > distsq
}
