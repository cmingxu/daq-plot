package main

type FourQuadPlot struct {
	*Plot
}

func NewQuadPlot() *FourQuadPlot {
	return &FourQuadPlot{}
}

func (p *FourQuadPlot) Update() {}
