package main

import (
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
	cfg    config
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	cfg := newConfig()
	app := &application{
		logger: logger,
		cfg:    cfg,
	}

	logger.Info("starting server", slog.String("addr", cfg.addr))
	logger.Debug("starting server", slog.String("addr", cfg.addr))

	err := http.ListenAndServe(cfg.addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
