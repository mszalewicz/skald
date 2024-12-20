package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"log/slog"
	"net/url"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mszalewicz/skald/assert"
	"github.com/mszalewicz/skald/database"
	"github.com/mszalewicz/skald/gui"
)

//go:embed schema.sql
var schema string
var backend database.Backend

const databaseName string = "skald.db"

func init() {
	{ // Create local db if it does not exist
		dbFile, err := os.OpenFile(databaseName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			slog.Error("Could not create log file.", "error", err)
		}
		defer dbFile.Close()
	}

	{ // Init database
		connectionUrlParams := make(url.Values)
		connectionUrlParams.Add("_txlock", "immediate")
		connectionUrlParams.Add("_journal_mode", "WAL")
		connectionUrlParams.Add("_busy_timeout", "5000")
		connectionUrlParams.Add("_synchronous", "NORMAL")
		connectionUrlParams.Add("_cache_size", "1000000000")
		connectionUrlParams.Add("_foreign_keys", "true")
		connectionUrl := "file:" + databaseName + "?" + connectionUrlParams.Encode()

		db, err := sql.Open("sqlite3", connectionUrl)

		if err != nil {
			log.Fatal(err)
		}

		backend.DB = db
		assert.NotNil(backend.DB)

		// Create tables
		if _, err := backend.DB.Exec(schema); err != nil {
			log.Fatal(err)
		}

		backend.Queries = database.New(backend.DB)
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

	localLog := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(localLog)

	{ // Get monitor resolution
		monitor := glfw.GetPrimaryMonitor()
		mode := monitor.GetVideoMode()
		screen.Width = int64(mode.Width)
		screen.Height = int64(mode.Height)
		width, height = gui.CalculateResolution(mode.Width, mode.Height, 0.9)
		minWidth, minHeight = gui.CalculateResolution(mode.Width, mode.Height, 0.5)

		glfw.Terminate()
	}

	{ // Read setting for given resolution
		fontsize, err := backend.Queries.GetFontSizeByWidth(context.Background(), screen.Width)

		if err != nil {
			log.Fatal(err)
		}

		if fontsize != 0 {
			settings.Fontsize = fontsize
		} else {
			settings.Fontsize = 10
		}
	}

	settings.Screen.Width = screen.Width
	settings.Screen.Height = screen.Height

	{ // Insert setting if it does not exist
		queries := database.New(backend.DB)
		settingOccurences, err := queries.CountSetting(context.Background(), settings.Screen.Width)

		if err != nil {
			log.Fatal(err)
		}

		if settingOccurences == 0 {
			err := queries.CreateSetting(context.Background(), database.CreateSettingParams{
				Width:    settings.Screen.Width,
				Height:   settings.Screen.Height,
				Fontsize: settings.Fontsize,
			})

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	go func() {
		window := new(app.Window)
		window.Option(app.Title("Skald"))
		window.Option(app.Size(unit.Dp(width), unit.Dp(height)))
		window.Option(app.MinSize(unit.Dp(minWidth), unit.Dp(minHeight)))

		assert.NotNil(backend)
		assert.NotNil(screen)
		assert.NotNil(settings)
		assert.NotNil(settings)
		assert.NotNil(window)

		{ // Check if local account exists
			numberOfAccounts, err := backend.Queries.CountAccounts(context.Background())

			if err != nil {
				log.Fatal(err)
			}

			_ = numberOfAccounts
			// if numberOfAccounts == 0 {
			// 	uuid := uuid.NewString()
			// 	fmt.Println(uuid)
			// 	err := gui.AccountCreation(window, &screen, &settings, &backend, uuid)

			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}
			// }
		}

		err := gui.MainWindow(window, &screen, &settings, &backend)

		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}()

	app.Main()
}
