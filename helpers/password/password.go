package password

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), 15)
	if err != nil {
		return "", fmt.Errorf("could not hash password", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPass string, candidatePass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(candidatePass))
}
