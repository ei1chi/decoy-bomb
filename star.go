package main

import (
	"math"
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

	blastR                   float64 // 炸裂半径
	rplsR                    float64 // 斥力半径
	rplsMag                  float64 // 斥力強さ
	count                    int
	pos, pPos, ppPos, pppPos complex128 // 位置、前回位置、前々回位置
	angle                    float64
	state                    StarStates
}

func (s *Star) update() bool {
	switch s.state {
	case starMoving:
		if !pressed { // 離した
			s.state = starFired
			s.count = 0
		} else {
			next := complex(cursorX, cursorY)
			d := next - s.pos
			if cmplx.Abs(d) > 0.5 {
				s.pppPos = s.ppPos
				s.ppPos = s.pPos
				s.pPos = s.pos
				s.pos = next
				b := s.pos - s.bezier()
				s.angle = math.Atan2(imag(b), real(b))
			}
			s.pos = next
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

func (s *Star) bezier() complex128 {
	t := math.Pow(0.5, 3)
	x := t*real(s.pppPos) + 3*t*real(s.ppPos) + 3*t*real(s.pPos) + t*real(s.pos)
	y := t*imag(s.pppPos) + 3*t*imag(s.ppPos) + 3*t*imag(s.pPos) + t*imag(s.pos)
	return complex(x, y)
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
	op.GeoM.Rotate(s.angle)
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
