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
	Screen   Screen
	Fontsize int
}

func MainWindow(window *app.Window, screen *Screen, settings *Settings) error {
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

				if ev.(key.Event).State == key.Press {
					name := ev.(key.Event).Name
					mod := ev.(key.Event).Modifiers

					if mod == key.ModCommand || mod == key.ModCtrl {
						switch name {
						case "=":
							if settings.Fontsize < 40 {
								settings.Fontsize += 1
							}
						case "-":
							if settings.Fontsize > 8 {
								settings.Fontsize -= 1

							}
						}

						fmt.Println(settings.Fontsize)
					}

				}
				// fmt.Printf("KEY: %+v\n", ev)
				// fmt.Println(settings.Fontsize)
			}

			e.Frame(gtx.Ops)

		}
	}
}
