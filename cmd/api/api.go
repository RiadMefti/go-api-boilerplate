package api

import (
	"log"
	"net/http"

	"github.com/RiadMefti/go-api-boilerplate/db"
)

type Server struct {
	address string
	store   db.Storage
}

func NewApiServer(adresse string, store db.Storage) *Server {

	return &Server{
		address: adresse,
		store:   store,
	}

}

func Run(s *Server) {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /login", login)
	mux.HandleFunc("POST /register", register)
	mux.HandleFunc("GET /protected/{id}", protectedRoute)
	log.Println("Server starting on", s.address, " ...")
	err := http.ListenAndServe(s.address, mux)
	if err != nil {
		log.Fatal(err)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))

}

func register(w http.ResponseWriter, r *http.Request) {

}
func protectedRoute(w http.ResponseWriter, r *http.Request) {

}
