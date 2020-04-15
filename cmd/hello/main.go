package main

import (
	"github.com/niclic/helloworld/internal/routing"
	"github.com/niclic/helloworld/internal/server"
)

func main() {
	server := server.HttpServer{8080, routing.DefaultRouter}
	server.Start()
}
