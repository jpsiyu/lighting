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

	priceRect := &Rect{x: 00, y: 0, x1: 80, y1: 20}
	price := NewPrice(priceRect)

	ui.Render(price.Widget())

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
