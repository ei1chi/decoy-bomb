package main

import (
	"math/cmplx"

	et "github.com/hajimehoshi/ebiten"
)

type Fragment struct {
	pos          complex128
	angle, speed float64
	count        int
	dead         bool
}

func (f *Fragment) update() {
	switch {
	case f.count == 0:
		f.speed = 3
	case f.count < 20:
	case f.count < 30:
		f.speed -= 0.3
	default:
		f.dead = true
	}
	f.pos += cmplx.Rect(f.speed, f.angle)
	f.count += 1
}

type StarStates int

const (
	noStar StarStates = iota
	starMoving
	starFired
	starBlasting
)

type Star struct {
	exist bool // for tank
	id    int  // for tank

	blastR  float64 // 炸裂半径
	rplsR   float64 // 斥力半径
	rplsMag float64 // 斥力強さ
	count   int
	pos     complex128
	state   StarStates
}

func (s *Star) update() bool {
	switch s.state {
	case starMoving:
		s.pos = complex(cursorX, cursorY)
		if !pressed { // 離した
			s.state = starFired
			s.count = 0
		}
	case starFired:
		if s.count > 64 {
			s.state = starBlasting
			s.count = 0
		}
	case starBlasting:
		if s.count > 20 {
			return false
		}
	}
	s.count += 1
	return true
}

func (s *Star) draw(sc *et.Image) {
	sp := sprites["star"]
	op := sp.center()
	switch s.state {
	case starFired:
		if s.count%8 < 4 {
			op.ColorM.Translate(0, 0, 0, -1)
		}
	}
	op.GeoM.Translate(real(s.pos), imag(s.pos))
	sc.DrawImage(sp.image, op)
}

// 型定義ここまで
//=========================

var (
	starsMax  = 3
	stars     = &Stars{}
	fragments = []Fragment{}
)

func initStars() {
	stars.init(starsMax)
}

func processStars() {

	// トリガーでスター生成
	// 押しっぱなしで移動・反発
	// 離して爆発
	if isJustPressed {
		s := Star{}
		s.pos = complex(cursorX, cursorY)
		s.state = starMoving
		stars.add(s)
	}

	for i, _ := range stars.arr {
		s := &stars.arr[i]
		if !s.exist {
			continue
		}
		if ok := s.update(); !ok {
			stars.remove(i)
		}
	}

	//for i, _ := range fragments {
	//f := &fragments[i]
	//f.update()
	//}
}
