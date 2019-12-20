package ui

import (
	"math"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Price struct{}

func NewPrice() *Price {
	return &Price{}
}

func (p *Price) Widget() *widgets.Plot {
	sinData := (func() []float64 {
		n := 220
		ps := make([]float64, n)
		for i := range ps {
			ps[i] = 1 + math.Sin(float64(i)/5)
		}
		return ps
	})()

	lc := widgets.NewPlot()
	lc.Title = "dot-marker Line Chart"
	lc.Data = make([][]float64, 1)
	lc.Data[0] = sinData
	lc.SetRect(0, 15, 50, 25)
	lc.AxesColor = ui.ColorWhite
	lc.LineColors[0] = ui.ColorRed
	lc.Marker = widgets.MarkerDot
	return lc
}
