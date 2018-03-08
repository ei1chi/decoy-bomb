// +build js

package main

import "github.com/gopherjs/gopherjs/js"

func getScale() float64 {
	doc := js.Global.Get("document")
	body := doc.Get("body")
	w := body.Get("clientWidth").Float()
	h := body.Get("clientHeight").Float()
	sw := w / 480
	sh := h / 640
	if sh < sw {
		return sh
	}
	return sw
}
