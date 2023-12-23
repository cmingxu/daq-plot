package main

import (
	"flag"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var windowOpt WindowOptions = WindowOptions{
	height:          800,
	width:           1400,
	statusBarHeight: 20,
	title:           "raylib [core] example - basic window",
	pos:             rl.Vector2{0, 0},
}

func init() {

	var (
		height = 800
		width  = 1400
	)
	flag.StringVar(&windowOpt.title, "title", "raylib [core] example - basic window", "window title")
	flag.IntVar(&height, "height", 800, "window height")
	flag.IntVar(&width, "width", 1400, "window width")
	flag.Parse()

	windowOpt.height = float32(height)
	windowOpt.width = float32(width)
}

func main() {
	w := NewDWindow(windowOpt)
	w.Show()
	w.Close()
}
