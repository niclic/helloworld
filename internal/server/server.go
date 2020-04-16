package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/niclic/helloworld/internal/routing"
)

// HTTPServer represents a basic HTTP Web Server.
type HTTPServer struct {
	Port   int
	router *routing.HTTPRouter
}

// SetRouter sets the HTTPRouter for this HTTPServer instance.
func (s *HTTPServer) SetRouter(r *routing.HTTPRouter) {
	s.router = r
}

// Start starts the HTTPServer on the specified localhost port.
func (s *HTTPServer) Start() {
	if s.router == nil {
		s.router = routing.NewHTTPRouter()
		s.router.RegisterRoutes(routing.Routes)
	}

	log.Printf("Listening on port %d...\n", s.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.router.Mux()))
}
