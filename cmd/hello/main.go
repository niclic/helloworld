package main

import (
	"fmt"
	"log"
	"net/http"
)

func createHttpRouter() *http.ServeMux {
	return http.NewServeMux()
}

func registerRoutes(routes map[string]httpHandler, router *http.ServeMux) {
	for r, h := range routes {
		router.HandleFunc(r, h)
	}
}

func startServer(port int, router *http.ServeMux) {
	log.Printf("Listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

type httpHandler func(w http.ResponseWriter, r *http.Request)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	port := 8080

	routes := map[string]httpHandler{
		"/": helloHandler,
	}

	router := createHttpRouter()
	registerRoutes(routes, router)
	startServer(port, router)
}
