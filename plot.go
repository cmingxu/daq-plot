package main

import (
	"log"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Plot struct {
	Updateable

	position rl.Vector2
	size     rl.Vector2

	padding   float32
	axisWidth float32

	title       string
	titleHeight float32

	bg *PlotBackground
}

func NewPlot() *Plot {
	return &Plot{
		titleHeight: 20,
		padding:     20,
		axisWidth:   20,
		bg:          NewPlotBackground(),
	}
}

func (p *Plot) SetPosition(position rl.Vector2) {
	p.position = position
	bgPosition := rl.Vector2{
		X: position.X + p.padding + p.axisWidth,
		Y: position.Y + p.titleHeight + p.padding,
	}
	p.bg.SetPosition(bgPosition)
}

func (p *Plot) SetSize(size rl.Vector2) {
	p.size = size
	bgSize := rl.Vector2{
		X: size.X - p.padding*2 - p.axisWidth,
		Y: size.Y - p.titleHeight - p.padding*2 - p.axisWidth,
	}
	p.bg.SetSize(bgSize)
}

func (p *Plot) SetTitle(title string) {
	p.title = title
}

func (p *Plot) Update() {
	log.Printf("Plot.Update()")
	gui.Panel(rl.NewRectangle(
		p.position.X,
		p.position.Y,
		p.size.X,
		p.size.Y,
	), p.title)
	p.bg.Update()
}
