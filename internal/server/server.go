package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/niclic/helloworld/internal/routing"
)

type HttpServer struct {
	Port   int
	router *routing.HttpRouter
}

func (s *HttpServer) SetRouter(r *routing.HttpRouter) {
	s.router = r
}

func (s *HttpServer) Start() {
	if s.router == nil {
		s.router = routing.DefaultRouter
	}

	log.Printf("Listening on port %d...\n", s.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.router))
}
