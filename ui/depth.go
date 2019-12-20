package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Depth struct {
	rect *Rect
}

func NewDepth(rect *Rect) *Depth {
	return &Depth{
		rect: rect,
	}
}

func simulate() []float64 {
	total := 80
	mid := 40
	array := make([]float64, total)
	index := 0
	for i := total; i > mid; i-- {
		array[index] = float64(i)
		index++
	}

	for i := 0; i < mid; i++ {
		array[index] = float64(i)
		index++
	}
	return array
}

func (d *Depth) Widget() *widgets.Plot {
	dc := widgets.NewPlot()
	dc.Title = "Depth Chart"
	dc.Data = make([][]float64, 1)
	dc.Data[0] = simulate()
	dc.SetRect(d.rect.x, d.rect.y, d.rect.x1, d.rect.y1)
	dc.AxesColor = ui.ColorCyan
	dc.LineColors[0] = ui.ColorBlue
	return dc
}
