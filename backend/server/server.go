package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

type Server struct {
	Router *chi.Mux
}

func StartServer() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		fmt.Println("PORT string is empty")
	}

	s := CreateServer()
	s.MountHandlers()

	http.ListenAndServe(":"+PORT, s.Router)
}

func CreateServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	s.Router.Use(middleware.Logger)
	s.Router.Use(corsMiddleware)
	return s
}

func (s *Server) MountHandlers() {
	s.Router.Get("/", serverGreeting)
	s.Router.Post("/add", sendSecretHandler)
	s.Router.Post("/get", getSecretHandler)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
