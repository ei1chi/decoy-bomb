package main

import (
	"math"
	"math/rand"

	et "github.com/hajimehoshi/ebiten"
)

type Decoy struct {
	exist bool // for tank
	id    int  // for tank

	power  int
	radius float64
	pos    complex128
}

type Decoys struct {
	h   TankHistory
	arr []Decoy
}

func (d *Decoys) init(size int) {
	d.h.Init(size)
	d.arr = make([]Decoy, size)
}

func (d *Decoys) add(val Decoy) error {
	idx, err := d.h.Pop()
	if err != nil {
		return err
	}
	d.arr[idx] = val
	return nil
}

func (d *Decoys) remove(idx int) {
	d.arr[idx].exist = false
	d.h.Push(idx)
}

const (
	decoysMax = 5
	powerMax  = 100
)

var (
	decoys      *Decoys
	isMouseHold bool
	nextDecoyId = 1
)

func initDecoys() {
	decoys = &Decoys{}
	decoys.init(decoysMax)
}

func processDecoys() {
	// トリガー処理
	x, y := et.CursorPosition()
	touched := false
	for _, t := range et.Touches() {
		x, y = t.Position()
		touched = true
	}
	if touched || et.IsMouseButtonPressed(et.MouseButtonLeft) {
		if !isMouseHold {
			d := Decoy{
				exist:  true,
				id:     nextDecoyId,
				power:  powerMax,
				radius: 125,
				pos:    complex(float64(x), float64(y)),
			}
			decoys.add(d)
			nextDecoyId += 1
		}
		isMouseHold = true
	} else {
		isMouseHold = false
	}

	// 減衰処理
	for i, _ := range decoys.arr {
		d := &decoys.arr[i]
		if !d.exist {
			continue
		}
		d.power -= 1
		if d.power < 0 {
			decoys.remove(i)
			for n := 0; n < 10; n++ {
				f := Fragment{}
				f.pos = d.pos
				f.angle = math.Pi*2*rand.Float64() - math.Pi
				fragments = append(fragments, f)
			}
		}
	}
}
