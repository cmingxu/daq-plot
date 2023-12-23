package main

import rl "github.com/gen2brain/raylib-go/raylib"

type LinePlot struct {
	plot *Plot
}

func NewLinePlot(position, size rl.Vector2, title string) *LinePlot {
	lp := &LinePlot{
		plot: NewPlot(),
	}

	lp.plot.SetPosition(position)
	lp.plot.SetSize(size)
	lp.plot.SetTitle(title)
	return lp
}

func (p *LinePlot) Update() {
	p.plot.Update()
}
