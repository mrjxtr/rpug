package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mrjxtr/rpug/internal/config"
	"github.com/mrjxtr/rpug/internal/data"
	"github.com/mrjxtr/rpug/internal/generator"
	"github.com/mrjxtr/rpug/internal/server"
)

//go:embed data/data.json
var dataJSON []byte

func main() {
	slog.Info("Loading configs...")
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Error loading config", "error", err)
		os.Exit(1)
	}

	slog.Info("Loading data...")
	var d data.Data
	if err := json.Unmarshal(dataJSON, &d); err != nil {
		slog.Error("Error loading data", "error", err)
		os.Exit(1)
	}

	slog.Info("Loading generators...")
	gen := generator.NewPinoyGenerator(cfg, &d)

	slog.Info("Loading server, middleware, and routes...")
	srv := server.NewServer(gen, cfg)

	httpSrv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           srv.SetupRouter(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	go func() {
		slog.Info(
			"Starting server",
			"port",
			cfg.Port,
			"ping_url",
			"http://localhost:3000/ping",
		)
		slog.Info(
			"Go to this url to test API",
			"api_url",
			"http://localhost:3000/api/v1/pinoys",
		)
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	slog.Info("Shutting down server")

	// Give the server 14 seconds to finish handling existing requests
	// NOTE: I've set the kill_timeout fly.toml and WriteTimeout to 15s
	shutdownCtx, shutdownCancel := context.WithTimeout(
		context.Background(),
		14*time.Second,
	)
	defer shutdownCancel()

	if err := httpSrv.Shutdown(shutdownCtx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	} else {
		slog.Info("Server gracefully stopped")
	}

	slog.Info("Cleanup completed")
}
