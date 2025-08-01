package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPwInBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPwInBytes), err
}
