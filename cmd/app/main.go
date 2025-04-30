package main

import (
	"log/slog"
	"os"

	"github.com/kafkaphoenix/gocube/internal/engine"
)

func main() {
	// Initialize the logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Initialize the engine
	app := engine.NewEngine(logger)

	// Start the engine
	if err := app.Start(); err != nil {
		panic(err)
	}
}
