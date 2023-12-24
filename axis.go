package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type AxisType uint8

const (
	XAxis AxisType = 1
	YAxis AxisType = 2
)

type TickPosition uint8

const (
	TickAbove TickPosition = 1
	TickBelow TickPosition = 2
	TickLeft  TickPosition = 3
	TickRight TickPosition = 4
)

type Axis struct {
	Updateable
	axisType     AxisType
	tickPosition TickPosition
	start        rl.Vector2
	end          rl.Vector2
	lineWidth    float32
	color        rl.Color

	title  string
	values []float32
}

func NewAxis(axisType AxisType, tickPosition TickPosition,
	title string, start, end rl.Vector2, values []float32) *Axis {
	return &Axis{
		axisType:     axisType,
		tickPosition: tickPosition,
		start:        start,
		end:          end,
		lineWidth:    2,
		color:        rl.Black,
		title:        title,
		values:       values,
	}
}

func (a *Axis) SetLineWidth(lineWidth float32) {
	a.lineWidth = lineWidth
}

func (a *Axis) Update() {
	a.draw()
}

func (a *Axis) draw() {
	rl.DrawLineEx(a.start, a.end, a.lineWidth, rl.Black)
	a.drawTicks()
}

func (a *Axis) drawTicks() {
	steps := len(a.values)
	switch a.axisType {
	case XAxis:
		for i := 0; i < steps; i++ {
			if a.tickPosition == TickBelow {
				x := a.start.X + float32(i)*(a.end.X-a.start.X)/float32(steps)
				rl.DrawLineEx(rl.Vector2{X: x, Y: a.start.Y}, rl.Vector2{X: x, Y: a.start.Y + 5}, a.lineWidth, rl.Black)
				rl.DrawText(fmt.Sprintf("%d", int(a.values[i])), int32(x), int32(a.start.Y+10), 10, rl.Black)
			} else if a.tickPosition == TickAbove {
				x := a.start.X + float32(i)*(a.end.X-a.start.X)/float32(steps)
				rl.DrawLineEx(rl.Vector2{X: x, Y: a.start.Y}, rl.Vector2{X: x, Y: a.start.Y - 5}, a.lineWidth, rl.Black)
				rl.DrawText(fmt.Sprintf("%d", int(a.values[i])), int32(x), int32(a.start.Y-10), 10, rl.Black)
			}
		}
	case YAxis:
		for i := 0; i < steps; i++ {
			if a.tickPosition == TickLeft {
				y := a.start.Y + float32(i)*(a.end.Y-a.start.Y)/float32(steps)
				rl.DrawLineEx(rl.Vector2{X: a.start.X, Y: y}, rl.Vector2{X: a.start.X - 10, Y: y}, a.lineWidth, rl.Black)
				rl.DrawText(fmt.Sprintf("%d", int(a.values[i])), int32(a.start.X-20), int32(y-5), 10, rl.Black)
			} else if a.tickPosition == TickRight {
				y := a.start.Y + float32(i)*(a.end.Y-a.start.Y)/float32(steps)
				rl.DrawLineEx(rl.Vector2{X: a.start.X, Y: y}, rl.Vector2{X: a.start.X + 5, Y: y}, a.lineWidth, rl.Black)
				rl.DrawText(fmt.Sprintf("%d", int(a.values[i])), int32(a.start.X+10), int32(y), 10, rl.Black)
			}
		}
	}
}
