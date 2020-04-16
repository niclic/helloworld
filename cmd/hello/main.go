package main

import (
	"github.com/niclic/helloworld/internal/server"
)

func main() {
	server := &server.HTTPServer{Port: 8080}
	server.Start()
}
