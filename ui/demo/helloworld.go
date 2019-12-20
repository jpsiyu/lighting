package demo

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type HelloWorld struct {
}

func NewHelloWorld() *HelloWorld {
	return &HelloWorld{}
}

func (hw *HelloWorld) Run() {
	if err := ui.Init(); err != nil {
		log.Fatalln("failed to initialize termui", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = "Hello World"
	p.SetRect(0, 0, 25, 5)
	ui.Render(p)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
