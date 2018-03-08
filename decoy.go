package main

import et "github.com/hajimehoshi/ebiten"

type Decoy struct {
	exist  bool
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
	powerMax  = 300
)

var (
	decoys      *Decoys
	isMouseHold bool
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
				power:  powerMax,
				radius: 125,
				pos:    complex(float64(x), float64(y)),
			}
			decoys.add(d)
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
		}
	}
}
