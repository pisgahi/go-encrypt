package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func serverGreeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello go-encrpt"))
}

func (s *Server) MountHandlers() {
	s.Router.Use(middleware.Logger)
	s.Router.Get("/", serverGreeting)
}
