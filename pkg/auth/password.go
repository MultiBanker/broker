package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const cost = 10 // сложность хэша, чем больше это число, тем дольше его подсчет

func HashPassword(password string) ([]byte, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}
	return hashed, nil
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


