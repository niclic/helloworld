package main

import (
	"fmt"
	"log"
	"net/http"
)

var mux = http.NewServeMux()

type httpRouter struct {
	*http.ServeMux
}

func (r *httpRouter) RegisterRoutes(routes map[string]httpHandler) {
	for p, h := range routes {
		r.HandleFunc(p, h)
	}
}

type httpServer struct {
	port   int
	router *httpRouter
}

func (s *httpServer) Start() {
	log.Printf("Listening on port %d...\n", s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router))
}

type httpHandler func(w http.ResponseWriter, r *http.Request)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	routes := map[string]httpHandler{
		"/hello": helloHandler,
	}

	router := &httpRouter{ServeMux: mux}
	router.RegisterRoutes(routes)

	server := &httpServer{8080, router}
	server.Start()
}
