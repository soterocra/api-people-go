package server

import (
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
func NewServer(port string, handler http.Handler) *Server {
	return &Server{
		port:    port,
		handler: handler,
	}
}

// Run é o método que "liga" o servidor
func (s *Server) Run() error {
	log.Printf("Servidor escutando em http://localhost%s", s.port)
	return http.ListenAndServe(s.port, s.handler)
}
