package server

import (
	"log/slog"
	"net/http"

	"github.com/mrjxtr/rpug/internal/views/layout"
	"github.com/mrjxtr/rpug/internal/views/pages"
)

// handlePinoysPage parses ?seed= and ?results= from the request, generates a
// deterministic PinoyResponse, and renders the playground page.
func (s *Server) handlePinoysPage(w http.ResponseWriter, r *http.Request) {
	results, err := s.getResultsParam(r)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			"invalid 'results' query parameter",
		)
		return
	}
	seed := getSeedParam(r)

	resp, err := s.gen.Generate(results, seed)
	if err != nil {
		slog.Error("generate failed", "error", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	if err := layout.Layout(pages.PinoysPage(resp, s.cfg.MaxResults), "RPUG | Playground").
		Render(r.Context(), w); err != nil {
		slog.Error("render failed", "error", err)
	}
}

// handleHomeRedirect bounces the root path to /pinoys with a 302.
// Placeholder until a real home page exists; 302 (not 301) so browsers don't cache it.
func (s *Server) handleHomeRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/pinoys", http.StatusFound)
}
