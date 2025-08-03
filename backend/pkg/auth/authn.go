package auth

import (
	"golang.org/x/crypto/bcrypt"
)


func Encrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// LIGHTME: In which case should I send error to upper layer?
	return string(hash), err
}


func Compare(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}