package main

import "math"

func collision() {
	for i, d := range decoys.arr {
		if !d.exist {
			continue
		}
		for _, g := range ghosts {
			gpos, gr := g.collInfo()
			diff := d.pos - gpos
			distsq := math.Pow(imag(diff), 2) + math.Pow(real(diff), 2)
			limitsq := math.Pow(d.radius+gr, 2)
			if limitsq > distsq {
				g.hit(i, &d)
			}
		}
	}
}
