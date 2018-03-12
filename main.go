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
	initStars()
}

func main() {
	loadSprites([]string{
		"galaxy",
		"ghost",
		"decoy",
		"circle",
		"castle",
		"quit",
		"star",
	})
	rand.Seed(time.Now().UnixNano())
	s := getScale()
	err := et.Run(update, 480, 640, s, "decoy and bomb")
	if err != nil && err != ErrSuccess {
		log.Fatal(err)
	}
}

func update(screen *et.Image) error {

	updateInput()

	processStars()
	processGenerator()
	processGalaxies()
	processGhosts()

	collision()

	drawAll(screen)
	sweepAll()

	// 終了判定
	quit := et.IsKeyPressed(et.KeyQ)
	if (480-32) < cursorX && cursorY < 32 {
		if pressed {
			quit = true
		}
	}
	if quit {
		return ErrSuccess
	}

	// FPS
	str := "FPS: %f\n"
	str += "1) hold and move the star. star pushes ghosts out\n"
	str += "2) release the star to delete ghosts around\n"
	str += "3) don't make ghosts arrive at your castle (<- not implemented)"
	ebitenutil.DebugPrint(screen, fmt.Sprintf(str, et.CurrentFPS()))

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
