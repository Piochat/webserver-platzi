package main

import (
	"log"
	"net/http"

	"gitlab.com/kokgudiel2/webserver/handlers"
	"gitlab.com/kokgudiel2/webserver/middleware"
	"gitlab.com/kokgudiel2/webserver/server"
)

func main() {
	server := server.NewServer(":8484")

	server.Handle(http.MethodGet, "/", server.AddMiddleware(handlers.HandleRoot, middleware.Logging()))
	server.Handle(http.MethodGet, "/home", server.AddMiddleware(handlers.HandleHome, middleware.CheckAuth(), middleware.Logging()))
	server.Handle("", "/**", handlers.HandleNotFound)
	server.Handle(http.MethodPost, "/user", server.AddMiddleware(handlers.PostRequest, middleware.Logging()))

	log.Fatalln(server.Listen())
}
