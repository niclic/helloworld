package main

import (
	"fmt"
	"log"
	"net/http"
)

func createHttpRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", helloHandler)
	return router
}

func startServer(port int, router *http.ServeMux) {
	log.Printf("Listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	defaultPort := 8080
	router := createHttpRouter()
	startServer(defaultPort, router)
}
