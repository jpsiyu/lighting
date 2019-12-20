package ui

import (
	"log"

	ui "github.com/gizak/termui/v3"
)

func Run() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	priceRect := &Rect{x: 0, y: 0, x1: 80, y1: 20}
	price := NewPrice(priceRect)

	depthRect := &Rect{x: 0, y: 20, x1: 80, y1: 35}
	depth := NewDepth(depthRect)

	ui.Render(price.Widget(), depth.Widget())

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
