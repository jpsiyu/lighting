package demo

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Paragraph struct{}

func NewParagraph() *Paragraph {
	return &Paragraph{}
}

func (p *Paragraph) Run() {
	if err := ui.Init(); err != nil {
		log.Fatal("failed to initialize termui", err)
	}
	defer ui.Close()

	p0 := widgets.NewParagraph()
	p0.Text = "Borderless Text"
	p0.SetRect(0, 0, 20, 5)
	p0.Border = false

	p1 := widgets.NewParagraph()
	p1.Title = "Label"
	p1.Text = "Hello World"
	p1.SetRect(20, 0, 35, 5)

	ui.Render(p0, p1)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
