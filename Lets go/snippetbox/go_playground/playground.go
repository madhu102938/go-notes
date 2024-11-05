package go_playground

import (
	"golang.org/x/crypto/bcrypt"
)

// hashing password
func HashPassword(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), 12)
}