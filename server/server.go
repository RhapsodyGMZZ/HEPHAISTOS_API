package server

import (
	m "hephaistos/middleware"
	"net/http"
)

type Server struct {
	Config http.Server
}

func (s *Server) HandleRoutes() {
	http.HandleFunc("/", m.DefaultPage)
}
