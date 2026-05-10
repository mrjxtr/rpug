package server

import "net/http"

// handleStaticFS serves the embedded static FS at /static/*.
// No StripPrefix needed: //go:embed keeps the "static/" prefix in the FS,
// which lines up with the URL prefix after http.FS strips the leading slash.
func (s *Server) handleStaticFS() http.Handler {
	return http.FileServer(http.FS(s.sfs))
}
