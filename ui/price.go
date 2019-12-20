package ui

import (
	"math"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Price struct {
	rect *Rect
}

func NewPrice(rect *Rect) *Price {
	return &Price{
		rect: rect,
	}
}

func sin(n int) []float64 {
	ps := make([]float64, n)
	for i := range ps {
		ps[i] = 1 + math.Sin(float64(i)/5)
	}
	return ps
}

func (p *Price) Widget() *widgets.Plot {
	sinData := sin(220)
	pc := widgets.NewPlot()
	pc.Title = "Price Chart"
	pc.Data = make([][]float64, 1)
	pc.Data[0] = sinData
	pc.SetRect(p.rect.x, p.rect.y, p.rect.x1, p.rect.y1)
	pc.AxesColor = ui.ColorWhite
	pc.LineColors[0] = ui.ColorBlue
	pc.Marker = widgets.MarkerDot
	return pc
}
