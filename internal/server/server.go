package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/niclic/helloworld/internal/routing"
)

type HttpServer struct {
	Port   int
	Router *routing.HttpRouter
}

func (s *HttpServer) Start() {
	log.Printf("Listening on port %d...\n", s.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.Router))
}
