package main

import (
	"log/slog"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mszalewicz/skald/gui"
	"github.com/mszalewicz/skald/network"
)

func init() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
}

func main() {
	var (
		screen    gui.Screen
		width     int
		height    int
		minWidth  int
		minHeight int
	)

	localLog := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(localLog)

	_ = network.Backend{}

	{ // Get monitor resolution
		monitor := glfw.GetPrimaryMonitor()
		mode := monitor.GetVideoMode()
		screen.Width = mode.Width
		screen.Height = mode.Height
		width, height = gui.CalculateResolution(mode.Width, mode.Height, 0.9)
		minWidth, minHeight = gui.CalculateResolution(mode.Width, mode.Height, 0.5)

		glfw.Terminate()
	}

	go func() {
		window := new(app.Window)
		window.Option(app.Title("Skald"))
		window.Option(app.Size(unit.Dp(width), unit.Dp(height)))
		window.Option(app.MinSize(unit.Dp(minWidth), unit.Dp(minHeight)))
		err := gui.MainWindow(window, &screen)

		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}()

	app.Main()
}
