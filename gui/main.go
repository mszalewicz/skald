package gui

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

type Screen struct {
	Width  int
	Height int
}

type Settings struct {
	screen   Screen
	fontsize int
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

			for {
				ev, ok := gtx.Event(
					key.Filter{Optional: key.ModCommand, Name: "="},
					key.Filter{Optional: key.ModCommand, Name: "-"},
				)

				if !ok {
					break
				}

				fmt.Printf("KEY: %+v\n", ev)
			}

			e.Frame(gtx.Ops)

		}
	}
}
