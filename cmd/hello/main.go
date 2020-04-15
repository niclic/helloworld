package main

import (
	"github.com/niclic/helloworld/internal/routing"
	"github.com/niclic/helloworld/internal/server"
)

func main() {
	routing.DefaultRouter.RegisterRoutes(routing.Routes)

	server := server.HttpServer{8080, routing.DefaultRouter}
	server.Start()
}
