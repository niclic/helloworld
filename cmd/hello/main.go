package main

import (
	"github.com/niclic/helloworld/internal/server"
)

func main() {
	server := &server.HttpServer{Port: 8080}
	server.Start()
}
