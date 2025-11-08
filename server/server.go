package server

import (
	"api-people-go/config"
	"log"
	"net/http"
)

type Server struct {
	port    string
	handler http.Handler
}

// NewServer é nossa "Fábrica"
// Note que pedimos um 'http.Handler' (uma INTERFACE)
// O nosso roteado '*http.ServeMux' implementa essa interface
func NewServer(cfg config.Config, handler http.Handler) *Server {
	return &Server{
		port:    cfg.Server.Port,
		handler: handler,
	}
}

// Run é o método que "liga" o servidor
func (s *Server) Run() error {
	log.Printf("Servidor escutando em http://localhost%s", s.port)
	return http.ListenAndServe(s.port, s.handler)
}
