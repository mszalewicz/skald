package gui

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

type Screen struct {
	Width  int
	Height int
}

func MainWindow(window *app.Window, screen *Screen) error {
	var ops op.Ops

	theme := material.NewTheme()
	theme.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			e.Frame(gtx.Ops)
		}
	}
}
