package gui

import (
	"context"
	"fmt"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"

	_ "embed"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/mszalewicz/skald/database"
)

type Screen struct {
	Width  int
	Height int
}

type Settings struct {
	Screen   Screen
	Fontsize int
}

func MainWindow(window *app.Window, screen *Screen, settings *Settings, backend *database.Backend) error {
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

						ctx := context.Background()
						queries := database.New(backend.DB)

						insertedSetting, err := queries.CreateSetting(ctx, database.CreateSettingParams{
							Width:    3840,
							Height:   2160,
							Fontsize: 16,
						})

						if err != nil {
							log.Fatal(err)
						}
						log.Println(insertedSetting)

						settings, err := queries.ListSettings(ctx)
						for _, setting := range settings {
							fmt.Println(setting.Width, setting.Height, setting.Fontsize)
						}

					}

				}
				// fmt.Printf("KEY: %+v\n", ev)
				// fmt.Println(settings.Fontsize)
			}

			e.Frame(gtx.Ops)

		}
	}
}
