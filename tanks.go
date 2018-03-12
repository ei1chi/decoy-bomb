package main

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
	b.arr[idx].exist = true
	b.arr[idx].id = b.nextId
	return nil
}

func (b *Stars) remove(idx int) {
	b.arr[idx].exist = false
	b.h.Push(idx)
}
