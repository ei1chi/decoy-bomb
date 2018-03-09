package main

import (
	"math/cmplx"
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

type StarStates int

const (
	starHold StarStates = iota
	starFired
	starBlasting
)

type Stars struct {
	h      TankHistory
	arr    []Star
	nextId int
}

func (b *Stars) init(size int) {
	b.h.Init(size)
	b.arr = make([]Star, size)
}

func (b *Stars) add(val Star) error {
	idx, err := b.h.Pop()
	if err != nil {
		return err
	}
	b.nextId += 1
	b.arr[idx] = val
	b.arr[idx].id = b.nextId
	return nil
}

func (b *Stars) remove(idx int) {
	b.arr[idx].exist = false
	b.h.Push(idx)
}

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

var (
	starsMax  = 3
	stars     = &Stars{}
	fragments = []Fragment{}
	//isMouseHold bool
)

func initStars() {
	stars.init(starsMax)
}

func processStars() {

	// 押しっぱなしで移動・反発
	for i, _ := range fragments {
		f := &fragments[i]
		f.update()
	}
}
