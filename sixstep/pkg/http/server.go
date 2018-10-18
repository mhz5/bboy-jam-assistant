package http

import (
	"bboy-jam-assistant/sixstep/cmd/sixstep"
)

type Server struct {
	router *Router
}

var _ sixstep.Server = &Server{}

func NewServer() sixstep.Server {
	return &Server{
		router: NewRouter(),
	}
}
func (s *Server) Serve() {
	s.router.Handle()
}
