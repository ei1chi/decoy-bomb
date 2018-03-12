package main

import (
	et "github.com/hajimehoshi/ebiten"
)

var (
	pressed          bool
	isJustPressed    bool
	isHold           bool
	cursorX, cursorY float64
)

func updateInput() {
	x, y := et.CursorPosition()
	pressed = et.IsMouseButtonPressed(et.MouseButtonLeft)
	for _, t := range et.Touches() {
		x, y = t.Position()
		if x+y > 0 {
			pressed = true
		}
	}
	cursorX, cursorY = float64(x), float64(y)

	if pressed {
		if !isHold {
			isJustPressed = true
		} else {
			isJustPressed = false
		}
		isHold = true
	} else {
		isHold = false
		isJustPressed = false
	}
}
