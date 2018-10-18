package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateSaltedPassword(password string) (string, error) {
	saltedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(saltedPassword), nil
}

func ComparePassword(expectedPass, actualPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(expectedPass), []byte(actualPass))
}
