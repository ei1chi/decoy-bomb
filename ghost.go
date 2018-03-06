package main

type Ghost struct {
	update func(*Ghost)
	pos    Point
	vec    Point
}

func normalGhost(g *Ghost) {
	g.pos.x += g.vec.x
	g.pos.y += g.vec.y
}

var (
	ghosts []*Ghost
)

func processGhosts() {

	for _, g := range ghosts {
		g.update(g)
	}
	// 画面外消滅処理
}
