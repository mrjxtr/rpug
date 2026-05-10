// Package server
package server

import (
	"io/fs"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/mrjxtr/rpug/internal/config"
	"github.com/mrjxtr/rpug/internal/generator"
)

const (
	rateLimitPerMinute   = 60
	gzipCompressionLevel = 5 // 1=fast, 9=best; middle ground
)

// Generator is the interface for generating Pinoy data.
type Generator interface {
	Generate(results int, seed string) (*generator.PinoyResponse, error)
}

type Server struct {
	gen Generator
	cfg *config.Config
	sfs fs.FS
}

// NewServer creates a new Server with a generator.
func NewServer(gen Generator, cfg *config.Config, sfs fs.FS) *Server {
	return &Server{
		gen: gen,
		cfg: cfg,
		sfs: sfs,
	}
}

// SetupRouter sets up the router for the server.
// returning a chi.Mux router with handlers, routes, and middleware.
func (s *Server) SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Heartbeat("/ping"))

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// ? NOTE: Rate limit per IP (~1 per second average)
	// ? prevents rapid-fire and forces the use of ?results=N for multiple items
	r.Use(httprate.LimitByRealIP(rateLimitPerMinute, time.Minute))

	r.Use(middleware.Compress(gzipCompressionLevel))

	r.Handle("/static/*", s.handleStaticFS())
	r.Get("/pinoys", s.handlePinoysPage)
	r.Get("/", s.handleHomeRedirect)

	r.Get("/api/v1/pinoys", s.handlePinoysAPI)

	return r
}
