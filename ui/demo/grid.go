package demo

import (
	"log"

	ui "github.com/gizak/termui/v3"
)

type Grid struct{}

func NewGrid() *Grid {
	return &Grid{}
}

func (g *Grid) Run() {
	if err := ui.Init(); err != nil {
		log.Fatalln("failed to initialize termui", err)
	}
	defer ui.Close()

}
