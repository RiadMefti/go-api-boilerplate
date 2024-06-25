package api

import (
	"log"
	"net/http"
)

type Server struct {
	address string
}

func NewApiServer(adresse string) *Server {

	return &Server{
		address: adresse,
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
