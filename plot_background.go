package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	GridLineDirectonHorizontal uint8 = iota
	GridLineDirectonVertical
)

type PlotBackground struct {
	Updateable

	position rl.Vector2
	size     rl.Vector2
	color    rl.Color

	gridLineDotSize     float32
	gridLineDotDistance int32
	gridSize            rl.Vector2
	gridLineColor       rl.Color
}

func NewPlotBackground() *PlotBackground {
	return &PlotBackground{
		color:               rl.White,
		gridSize:            rl.Vector2{X: 30, Y: 30},
		gridLineDotSize:     2,
		gridLineDotDistance: 10,
		gridLineColor:       rl.DarkGray,
	}
}

func (p *PlotBackground) SetColor(color rl.Color) {
	p.color = color
}

func (p *PlotBackground) SetPosition(position rl.Vector2) {
	p.position = position
}

func (p *PlotBackground) SetSize(size rl.Vector2) {
	p.size = size
}

func (p *PlotBackground) SetGridSize(gridSize rl.Vector2) {
	p.gridSize = gridSize
}

func (p *PlotBackground) SetGridLineColor(gridLineColor rl.Color) {
	p.gridLineColor = gridLineColor
}

func (p *PlotBackground) Update() {
	p.draw()
}

func (p *PlotBackground) draw() {
	p.drawGrid()
}

func (p *PlotBackground) drawGrid() {
	var mouseCell rl.Vector2
	gui.Grid(rl.Rectangle{
		X:      p.position.X,
		Y:      p.position.Y,
		Width:  p.size.X,
		Height: p.size.Y,
	}, "", 50, 2, &mouseCell)
}
