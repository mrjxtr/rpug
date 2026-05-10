package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

// respondWithJSON encodes payload as JSON and writes it with the given status code.
func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		slog.Error("Error encoding json", "error", err)
		return
	}
}

// respondWithError writes a JSON error body of the form {"error": msg} with the given status code.
func respondWithError(w http.ResponseWriter, code int, msg string) {
	type respBody struct {
		Error string `json:"error"`
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(respBody{Error: msg}); err != nil {
		slog.Error("Error encoding json", "error", err)
	}
}

// getResultsParam parses ?results=n from the request and returns the number of results.
// defaulting to 1 and clamping to the provided max. Returns an error if the value is not an integer.
func (s *Server) getResultsParam(r *http.Request) (int, error) {
	results := r.URL.Query().Get("results")
	if results == "" {
		return 1, nil
	}

	resultsInt, err := strconv.Atoi(results)
	if err != nil {
		return 0, err
	}

	if resultsInt < 1 {
		resultsInt = 1
	}

	if resultsInt > s.cfg.MaxResults {
		resultsInt = s.cfg.MaxResults
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
