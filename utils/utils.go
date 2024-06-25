package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassowrd(password string) (string, error) {

	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(encpw), nil
}

func ValidatePassoword(password string, encryptedPassowrd string) bool {

	return bcrypt.CompareHashAndPassword([]byte(encryptedPassowrd), []byte(password)) == nil

}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

var jwtKey = []byte(os.Getenv("JWTSECRET"))

type CustomClaims struct {
	UserID string `json:"id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new JWT token for a user.
func GenerateToken(userID, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := CustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "api-boilerplate",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the JWT token and returns the custom claims.
func ValidateToken(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
func GenerateRandomID() (int, error) {
	// Define the maximum value for the ID as the maximum for a 32-bit signed integer
	max := big.NewInt(2147483647) // Maximum value for PostgreSQL integer type
	// Generate a cryptographically secure random big.Int less than max
	id, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	// Ensure the value is positive
	if id.Sign() == -1 {
		id = id.Neg(id)
	}
	// Convert the big.Int to int and return
	return int(id.Int64()), nil
}
