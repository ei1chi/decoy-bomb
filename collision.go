package main

import "math"

func collision(d *Decoy, g Enemy) {
	gx, gy, gr := g.collInfo()
	distsq := math.Pow((d.pos.x-gx), 2) + math.Pow((d.pos.y-gy), 2)
	limitsq := math.Pow((d.radius + gr), 2)
	if limitsq > distsq {
		g.hit(1)
		//vx, vy := (d.pos.x-g.pos().x)/distsq, (d.pos.y-g.pos().y)/distsq
		//g.vec.x += vx * 0.7
		//g.vec.y += vy * 0.7
	}
}
