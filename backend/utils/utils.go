package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash for password: %w", err)
	}

	return string(bytes), nil
}
