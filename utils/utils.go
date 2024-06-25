package utils

import "golang.org/x/crypto/bcrypt"

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
