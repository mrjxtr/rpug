// Package server
package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/mrjxtr/rpug/internal/config"
	"github.com/mrjxtr/rpug/internal/generator"
)

type Server struct {
	generator generator.Generator
	cfg       *config.Config
}

// NewServer creates a new Server with a generator.
func NewServer(gen generator.Generator, cfg *config.Config) *Server {
	return &Server{
		generator: gen,
		cfg:       cfg,
	}
}

// SetupRouter sets up the router for the server.
// returning a chi.Mux router with handlers, routes, and middleware.
func (s *Server) SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// ? NOTE: Rate limit to 2 requests per second per IP
	// ? prevents rapid-fire and forces the use of ?results=N for multiple items
	r.Use(httprate.LimitByRealIP(2, time.Second))

	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Compress(5))

	//? NOTE: Could be refactor into a handlers package in the future
	//? But this will be good enough for now
	// Handler for generating random Pinoy users.
	r.Get("/api/v1/pinoys", func(w http.ResponseWriter, r *http.Request) {
		// NOTE: If seed is present, generate data based on seed
		seedParam := getSeedParam(r)

		resParam, err := getResultsParam(r)
		if err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusBadRequest),
				http.StatusBadRequest,
			)
			return
		}

		resp, err := s.generator.Generate(resParam, seedParam)
		if err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		writeJSON(w, resp)
	})

	return r
}

// writeJSON writes data to the ResponseWriter as JSON. Simple and chill.
func writeJSON(w http.ResponseWriter, data any) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("Encode failed", "error", err)
	}
}

// getResultsParam parses ?results=n from the request and returns the number of results.
// defaulting to 1 and clamping to 1000. Returns an error if the value is not an integer.
func getResultsParam(r *http.Request) (int, error) {
	results := r.URL.Query().Get("results")
	if results == "" {
		return 1, nil
	}

	resultsInt, err := strconv.Atoi(results)
	if err != nil {
		return 1, err
	}

	if resultsInt < 1 {
		resultsInt = 1
	}

	if resultsInt > 1000 {
		resultsInt = 1000
	}

	return resultsInt, nil
}

// getSeedParam parses ?seed= from the request and returns the string of seed.
func getSeedParam(r *http.Request) string {
	seed := r.URL.Query().Get("seed")
	if seed != "" {
		return seed
	}
	return ""
}
