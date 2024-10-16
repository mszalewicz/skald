package main

import (
	"database/sql"
	_ "embed"
	"log"
	"log/slog"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mszalewicz/skald/database"
	"github.com/mszalewicz/skald/gui"
	"github.com/mszalewicz/skald/network"
)

//go:embed schema.sql
var schema string
var backend database.Backend

func init() {

	{ // Create local db if it does not exist
		dbFile, err := os.OpenFile("db.sqlite", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			slog.Error("Could not create log file.", "error", err)
		}
		defer dbFile.Close()
	}

	{ // Init database
		db, err := sql.Open("sqlite3", "db.sqlite")

		if err != nil {
			log.Fatal(err)
		}

		backend.DB = db

		// Create tables
		if _, err := backend.DB.Exec(schema); err != nil {
			log.Fatal(err)
		}
	}

	if err := glfw.Init(); err != nil {
		panic(err)
	}
}

func main() {
	var (
		settings  gui.Settings
		screen    gui.Screen
		width     int
		height    int
		minWidth  int
		minHeight int
	)

	settings.Fontsize = 10

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
		err := gui.MainWindow(window, &screen, &settings, &backend)

		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}()

	app.Main()
}
