package main

import "math"

func collision(d *Decoy, g *Ghost) {
	distsq := math.Pow((d.pos.x-g.pos.x), 2) + math.Pow((d.pos.y-g.pos.y), 2)
	limitsq := math.Pow((d.radius + 16), 2)
	if limitsq > distsq {
		vx, vy := (d.pos.x-g.pos.x)/distsq, (d.pos.y-g.pos.y)/distsq
		g.vec.x += vx * 0.7
		g.vec.y += vy * 0.7
	}
}
