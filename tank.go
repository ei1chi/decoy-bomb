package main

import (
	"errors"
	"log"
)

var (
	ErrTankEmpty = errors.New("tank is empty when pop()")
	ErrTankFull  = errors.New("tank is full when push()")
)

type TankHistory struct {
	queue      []int
	begin, end int
	filled     bool
}

func (h *TankHistory) Init(size int) {
	h.queue = make([]int, size)
	for i, _ := range h.queue {
		h.queue[i] = i
	}
	h.begin = 0
	h.end = 0
	h.filled = true
}

func (h *TankHistory) Push(val int) {
	if h.filled {
		log.Fatal(ErrTankFull)
	}
	h.queue[h.end] = val
	h.end += 1
	if h.end == len(h.queue) {
		h.end = 0 // loop
	}
	if h.end == h.begin {
		h.filled = true
	}
}

func (h *TankHistory) Pop() (int, error) {
	if h.begin == h.end && !h.filled {
		return 0, ErrTankEmpty
	}
	h.filled = false
	val := h.queue[h.begin]
	h.begin += 1
	if h.begin == len(h.queue) {
		h.begin = 0 // loop
	}
	return val, nil
}
