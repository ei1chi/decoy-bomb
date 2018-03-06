package main

import (
	"fmt"
	"log"

	et "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Sprite struct {
	image        *et.Image
	halfW, halfH float64
}

func newSprite(path string) (*Sprite, error) {
	s := &Sprite{}
	var err error
	s.image, _, err = ebitenutil.NewImageFromFile(path, et.FilterDefault)
	w, h := s.image.Size()
	s.halfW = float64(w) / 2
	s.halfH = float64(h) / 2
	return s, err
}

func (s *Sprite) center() *et.DrawImageOptions {
	op := &et.DrawImageOptions{}
	op.GeoM.Translate(-s.halfW, -s.halfH)
	return op
}

var sprites = map[string]*Sprite{}

func loadSprites(pngs []string) {
	var err error
	for _, name := range pngs {
		path := fmt.Sprintf("resources/%s.png", name)
		sprites[name], err = newSprite(path)
		if err != nil {
			log.Fatal(err)
		}
	}
}
