package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/RiadMefti/go-api-boilerplate/db"
	"github.com/RiadMefti/go-api-boilerplate/types"
	"github.com/RiadMefti/go-api-boilerplate/utils"
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
	mux.HandleFunc("POST /login", s.login)
	mux.HandleFunc("POST /register", s.register)
	mux.HandleFunc("GET /protected/{id}", s.protectedRoute)
	log.Println("Server starting on", s.address, "...")
	if err := http.ListenAndServe(s.address, mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}

}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	var userRq types.LoginUserRq

	errBody := utils.ParseJSON(r, &userRq)
	if errBody != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid request body"))

		return
	}
	user, err := s.store.GetUserByEmail(r.Context(), userRq.Email)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("internal server error"))
		return
	}

	samePassword := utils.ValidatePassword(userRq.Password, user.EncryptedPassword)
	if !samePassword {
		utils.WriteError(w, http.StatusUnauthorized, errors.New("invalid email or password"))
		return
	}

	jwt, jwtErr := utils.GenerateToken(fmt.Sprint(user.ID), userRq.Email)
	if jwtErr != nil {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("something went wrong"))

		return
	}

	utils.WriteJSON(w, 200, struct {
		Jwt string `json:"jwt"`
	}{Jwt: jwt})

}

func (s *Server) register(w http.ResponseWriter, r *http.Request) {

	var userRq types.RegisterUserRq

	errBody := utils.ParseJSON(r, &userRq)
	if errBody != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}

	hashedPassowrd, err := utils.HashPassword(userRq.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		return
	}

	id, err := utils.GenerateRandomID()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		return
	}

	newUser := types.User{
		ID:                id,
		EncryptedPassword: hashedPassowrd,
		Username:          userRq.Username,
		Email:             userRq.Email,
	}

	err = s.store.CreateUser(r.Context(), &newUser)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	jwt, jwtErr := utils.GenerateToken(fmt.Sprint(id), userRq.Email)
	if jwtErr != nil {
		utils.WriteError(w, 500, err)
		return
	}

	utils.WriteJSON(w, 200, struct {
		Jwt string `json:"jwt"`
	}{Jwt: jwt})

}
func (s *Server) protectedRoute(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	id := r.PathValue("id")
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if claims.UserID != id {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Proceed with the protected resource access
	response := fmt.Sprintf("Access granted to user %s with email %s", claims.UserID, claims.Email)
	w.Write([]byte(response))

}
