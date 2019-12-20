package ui

import (
	"log"

	"github.com/jpsiyu/lighting/ui/demo"
)

func Run() {
	log.Println("UI system running...")

	ins := demo.NewParagraph()
	ins.Run()
}
