package server

import "net/http"

// handlePinoysAPI parses ?seed= and ?results= from the request, generates a
// deterministic PinoyResponse, and writes it as JSON.
func (s *Server) handlePinoysAPI(w http.ResponseWriter, r *http.Request) {
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
		respondWithError(
			w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
		return
	}

	respondWithJSON(w, http.StatusOK, resp)
}
