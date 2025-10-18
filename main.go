package main

import (
	_ "embed"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

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
	http.ListenAndServe(":"+cfg.Port, srv.SetupRouter())
}
