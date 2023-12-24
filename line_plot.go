package main

import rl "github.com/gen2brain/raylib-go/raylib"

type LinePlot struct {
	base  *Plot
	xAxis *Axis
	yAxis *Axis
}

func NewLinePlot(position, size rl.Vector2, title string) *LinePlot {
	lp := &LinePlot{
		base: NewPlot(),
		xAxis: NewAxis(
			XAxis,
			TickBelow,
			"X Axis",
			rl.Vector2{X: position.X + 20 + 20, Y: position.Y + size.Y - 20 - 20},
			rl.Vector2{X: position.X + size.X - 20, Y: position.Y + size.Y - 20 - 20},
			[]float32{0, 1, 2, 3, 4, 5},
		),

		yAxis: NewAxis(
			YAxis,
			TickLeft,
			"Y Axis",
			rl.Vector2{X: position.X + 20 + 20, Y: position.Y + size.Y - 20 - 20},
			rl.Vector2{X: position.X + 20 + 20, Y: position.Y + 20 + 20},
			[]float32{0, 1, 2, 3, 4, 9},
		),
	}

	lp.base.SetPosition(position)
	lp.base.SetSize(size)
	lp.base.SetTitle(title)
	return lp
}

func (p *LinePlot) Update() {
	p.base.Update()
	p.xAxis.Update()
	p.yAxis.Update()
}
