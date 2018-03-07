package main

import (
	"errors"
	"math/rand"
	"time"

	et "github.com/hajimehoshi/ebiten"
)

type Point struct {
	x, y float64
}

func init() {
	loadSprites([]string{
		"galaxy",
		"ghost",
		"decoy",
		"circle",
		"castle",
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	et.Run(update, 480, 640, 1, "decoy and bomb")
}

func update(screen *et.Image) error {

	processGenerator()
	processDecoys()
	processGalaxies()
	processGhosts()

	// 当たり判定とベクトル変化
	for _, d := range decoys {
		for _, g := range ghosts {
			collision(d, g)
		}
	}

	drawAll(screen)
	sweepAll()

	quit := et.IsKeyPressed(et.KeyQ)
	if quit {
		return errors.New("success")
	}

	return nil
}

func sweepAll() {
	ngx := galaxies[:0]
	for _, g := range galaxies {
		if !g.isDead() {
			ngx = append(ngx, g)
		}
	}
	galaxies = ngx

	ngs := ghosts[:0]
	for _, g := range ghosts {
		if !g.isDead() {
			ngs = append(ngs, g)
		}
	}
	ghosts = ngs
}
