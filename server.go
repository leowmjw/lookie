package lookie

import "net/http"

// Server object
// TODO this is a stub
type Server struct {
	*http.Server
}

func (s *Server) Start() error {
	return s.ListenAndServe()
}
