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

	price := NewPrice()

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
