package gui

import (
	"context"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"

	_ "embed"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/mszalewicz/skald/assert"
	"github.com/mszalewicz/skald/database"
)

type Screen struct {
	Width  int64
	Height int64
}

type Settings struct {
	Screen   Screen
	Fontsize int64
}

func MainWindow(
	window *app.Window,
	screen *Screen,
	settings *Settings,
	backend *database.Backend,
) error {
	assert.NotNil(window)
	assert.NotNil(screen)
	assert.NotNil(settings)
	assert.NotNil(backend)

	var ops op.Ops

	theme := material.NewTheme()
	theme.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	text := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec vel nisl vestibulum purus interdum sollicitudin. Pellentesque sodales velit eu odio varius euismod. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed eleifend nisl eget enim cursus posuere. Nunc felis arcu, rutrum vitae vulputate sit amet, porttitor scelerisque eros. Maecenas in lorem eu magna venenatis eleifend. Etiam tincidunt elit non tincidunt semper. Curabitur id porttitor mauris. Praesent euismod quam ut leo ullamcorper, sed sagittis ante finibus. Nullam ut sollicitudin quam, lacinia interdum nisi. Nunc volutpat velit venenatis, gravida lorem in, feugiat turpis.

    Integer sem enim, elementum vel dui in, fringilla euismod sem. Maecenas dignissim, mauris sit amet feugiat molestie, elit urna tincidunt urna, nec lobortis odio ex a ligula. Etiam pellentesque lorem a est venenatis, a tincidunt mauris commodo. Sed id porttitor nisl. Fusce eleifend posuere odio at consequat. Nullam consectetur nisi eget lorem posuere, id auctor enim scelerisque. Sed ultrices aliquet imperdiet. Nullam varius mauris ac pharetra hendrerit. Curabitur egestas nulla ut dolor posuere dictum. Nulla facilisi. Nulla in venenatis tellus, et rutrum ante.

    Etiam scelerisque mattis massa, quis blandit nulla blandit vitae. Duis neque est, cursus nec metus vitae, lacinia tristique sem. Fusce rutrum scelerisque risus, eget tempor turpis blandit a. Fusce scelerisque ante a metus luctus, id ornare nulla hendrerit. Nam sed lacinia lorem, eget tempus dui. Donec sed finibus sapien. Vivamus blandit libero commodo, posuere erat vitae, lobortis tortor.

    Nulla blandit mauris in porttitor imperdiet. Sed gravida lectus varius convallis vehicula. Praesent viverra molestie nulla, non mollis tortor porta nec. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis arcu neque, molestie vitae est vel, congue hendrerit orci. Ut tincidunt orci ultricies luctus lobortis. Nunc sed justo id arcu dignissim vehicula. Mauris eget vulputate ex. Praesent hendrerit massa a dolor porta tempus. Vestibulum ultrices lacus et massa luctus lacinia. Nunc semper tempus diam vel posuere. Curabitur nec cursus augue.

    Nam ac efficitur metus. Interdum et malesuada fames ac ante ipsum primis in faucibus. Phasellus sodales neque est, eu hendrerit lacus viverra eu. Integer ipsum lectus, congue non lobortis id, vestibulum ac augue. Nam ex libero, semper sit amet justo quis, consequat imperdiet ipsum. In ac tempus nibh. Maecenas non sapien id urna tincidunt pretium sit amet in tellus. Mauris sagittis lectus rhoncus, aliquam quam et, suscipit erat. Aenean feugiat sit amet ante sit amet mollis. Quisque rhoncus viverra tempor. Pellentesque id ipsum imperdiet quam convallis interdum. Nunc condimentum lectus ut varius pellentesque.
    `

	list := widget.List{List: layout.List{Axis: layout.Vertical, Alignment: layout.Start}}

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			MainWindowWidget(&gtx, &list, theme, &text, settings)

			for {
				ev, ok := gtx.Event(
					key.Filter{Optional: key.ModCommand, Name: "="},
					key.Filter{Optional: key.ModCommand, Name: "-"},
				)

				// fmt.Printf("KEY: %+v\n", ev)

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

						queries := database.New(backend.DB)

						settingOccurences, err := queries.CountSetting(
							context.Background(),
							int64(settings.Screen.Width),
						)

						if err != nil {
							log.Fatal(err)
						}

						if settingOccurences == 0 {
							err := queries.CreateSetting(
								context.Background(),
								database.CreateSettingParams{
									Width:    settings.Screen.Width,
									Height:   settings.Screen.Height,
									Fontsize: settings.Fontsize,
								})

							if err != nil {
								log.Fatal(err)
							}
						} else {
							err := queries.UpdateSettingFont(
								context.Background(),
								database.UpdateSettingFontParams{
									Fontsize: settings.Fontsize,
									Width:    settings.Screen.Width,
								})

							if err != nil {
								log.Fatal(err)
							}
						}

						go func() {
							window.Invalidate()
						}()
					}
				}
			}

			e.Frame(gtx.Ops)

		}
	}
}

func AccountCreation(
	window *app.Window,
	screen *Screen,
	settings *Settings,
	backend *database.Backend,
	uuid string,
) error {
	assert.NotNil(window)
	assert.NotNil(screen)
	assert.NotNil(settings)
	assert.NotNil(backend)
	assert.Assert(len(uuid) == 36)

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
