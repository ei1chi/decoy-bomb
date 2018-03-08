package main

import (
	"errors"
	"fmt"
	"log"
	"math/cmplx"
	"math/rand"
	"time"

	et "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	ErrSuccess = errors.New("successfully finished")
)

const (
	around = 4.0 // 4 phases
)

func powi(angle float64) complex128 {
	return cmplx.Pow(1i, complex(angle, 0))
}

func init() {
	initDecoys()
}

func main() {
	loadSprites([]string{
		"galaxy",
		"ghost",
		"decoy",
		"circle",
		"castle",
		"quit",
	})
	rand.Seed(time.Now().UnixNano())
	err := et.Run(update, 480, 640, 1, "decoy and bomb")
	if err != nil && err != ErrSuccess {
		log.Fatal(err)
	}
}

func update(screen *et.Image) error {

	processGenerator()
	processDecoys()
	processGalaxies()
	processGhosts()

	// 当たり判定とベクトル変化
	collision()

	drawAll(screen)
	sweepAll()

	// 終了判定
	quit := et.IsKeyPressed(et.KeyQ)
	x, y := et.CursorPosition()
	touched := false
	for _, t := range et.Touches() {
		x, y = t.Position()
		touched = true
	}
	if (480-32) < x && y < 32 {
		if touched || et.IsMouseButtonPressed(et.MouseButtonLeft) {
			quit = true
		}
	}
	if quit {
		return errors.New("success")
	}

	// FPSほかデバッグ用
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f\nGhosts: %d", et.CurrentFPS(), len(ghosts)))

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
