// Package server
package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrjxtr/rpug/internal/generator"
)

type Server struct {
	gen generator.Generator
}

// NewServer creates a new Server.
func NewServer(gen generator.Generator) *Server {
	return &Server{
		gen: gen,
	}
}

// SetupRouter sets up the router for the server.
// returning a chi.Mux router with handlers, routes, and middleware.
func (s *Server) SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Compress(5))

	//? NOTE: Could be refactor into a handlers package in the future
	//? But this will be good enough for now
	// Handler for generating random Pinoy users.
	r.Get("/api/v1/pinoys", func(w http.ResponseWriter, r *http.Request) {
		results, err := getResults(r)
		if err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusBadRequest),
				http.StatusBadRequest,
			)
			return
		}

		resp, err := s.gen.Generate(results)
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
		slog.Error("encode failed", "error", err)
	}
}

// getResults parses ?results=n from the request and returns the number of results.
// defaulting to 1 and clamping to 1000. Returns an error if the value is not an integer.
func getResults(r *http.Request) (int, error) {
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
