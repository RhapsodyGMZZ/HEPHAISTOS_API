package server

import (
	routes "hephaistos/server/routes"
	"net/http"
)

type Server struct {
	Config http.Server
}

func (s *Server) HandleRoutes() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", routes.Home)
	http.HandleFunc("/FDS", routes.FDS)
}
