package main

import (
	"log"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type WindowOptions struct {
	height float32
	width  float32
	title  string
	color  rl.Color
	pos    rl.Vector2

	statusBarHeight float32
}

type DWindow struct {
	Updateable

	widgetsPanel *WidgetsPanel
	linePlot     *LinePlot
	fourQuadPlot *FourQuadPlot
	opts         *WindowOptions

	statusBarText string
	updateables   []Updateable
}

func NewDWindow(opts WindowOptions) *DWindow {
	w := DWindow{}
	w.opts = &opts
	w.widgetsPanel = NewWidgetsPanel()
	w.linePlot = NewLinePlot(
		rl.Vector2{X: 0, Y: 0},
		rl.Vector2{X: float32(w.opts.width * 3 / 4), Y: float32(w.opts.height - w.opts.statusBarHeight)},
		"Line Plot",
	)
	w.fourQuadPlot = NewQuadPlot()
	w.statusBarText = "This is a status bar"

	w.updateables = []Updateable{
		w.widgetsPanel,
		w.linePlot,
		w.fourQuadPlot,
	}
	return &w
}

func (w *DWindow) Show() {
	rl.InitWindow(int32(w.opts.width),
		int32(w.opts.height),
		w.opts.title)

	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowHighdpi | rl.FlagWindowUndecorated)
	rl.SetTargetFPS(60)
	rl.SetExitKey(0)

	w.LogStats()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		gui.StatusBar(rl.Rectangle{0, float32(rl.GetScreenHeight()) - w.opts.statusBarHeight,
			float32(rl.GetScreenWidth()), w.opts.statusBarHeight}, w.statusBarText)
		w.Update()

		rl.EndDrawing()
	}
}

func (w *DWindow) Update() {
	log.Printf("Window update")

	for _, u := range w.updateables {
		u.Update()
	}
	w.handleKeyevent()
	w.handleMouseevent()
}

func (w *DWindow) LogStats() {
	log.Printf("Window stats: %v", rl.GetWindowPosition())
}

func (w *DWindow) Close() {
	rl.CloseWindow()
}

func (w *DWindow) handleKeyevent() {
	if rl.IsKeyPressed(rl.KeySpace) {
		log.Printf("Space pressed")
	}
}

func (w *DWindow) handleMouseevent() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		log.Printf("Mouse left button pressed")
	}
}
