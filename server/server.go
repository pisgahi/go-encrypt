package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
}

func StartServer() {
	s := CreateServer()
	s.MountHandlers()

	http.ListenAndServe(":8080", s.Router)
}

func CreateServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountHandlers() {
	// s.Router.Use(middleware.Logger)
	s.Router.Get("/", serverGreeting)
	s.Router.Post("/add", sendSecret)
	s.Router.Get("/get", getSecretHandler)
}
