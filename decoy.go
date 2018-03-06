package main

import et "github.com/hajimehoshi/ebiten"

type Decoy struct {
	power  int
	radius float64
	pos    Point
}

func newDecoy(pos Point) *Decoy {
	return &Decoy{
		power:  powerMax,
		radius: 125,
		pos:    pos,
	}
}

const (
	decoysMax = 3
	powerMax  = 300
)

var (
	decoys      []*Decoy
	isMouseHold bool
)

func processDecoys() {
	// トリガー処理
	x, y := et.CursorPosition()
	if et.IsMouseButtonPressed(et.MouseButtonLeft) {
		if !isMouseHold && len(decoys) < decoysMax {
			decoys = append(decoys, newDecoy(Point{float64(x), float64(y)}))
		}
		isMouseHold = true
	} else {
		isMouseHold = false
	}

	// 減衰処理
	next := make([]*Decoy, 0, decoysMax)
	for _, d := range decoys {
		d.power -= 1
		if d.power > 0 {
			next = append(next, d)
		}
	}
	decoys = next
}
