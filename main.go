package main

import (
	"errors"

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
	})
}

func main() {
	et.Run(update, 480, 640, 1, "decoy and bomb")
}

func update(screen *et.Image) error {

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

	quit := et.IsKeyPressed(et.KeyQ)
	if quit {
		return errors.New("success")
	}

	return nil
}
