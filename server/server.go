package server

import (
	"log"
	"net/http"

	"gitlab.com/kokgudiel2/webserver/middleware"
	"gitlab.com/kokgudiel2/webserver/router"
)

type Server struct {
	port   string
	router *router.Router
}

func NewServer(port string) *Server {

	return &Server{port: port, router: router.NewRouter()}
}

func (s *Server) Handle(method, path string, handler http.HandlerFunc) {
	_, exist := s.router.Rules[path]
	if !exist {
		s.router.Rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.Rules[path][method] = handler
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...middleware.Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}

func (s *Server) Listen() error {
	log.Println("Server...")

	http.Handle("/", s.router)

	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}

	return nil
}
