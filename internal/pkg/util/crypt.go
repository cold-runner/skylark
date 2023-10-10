package util

import (
	"golang.org/x/crypto/bcrypt"
)

// Crypt Encrypt the password using crypto/bcrypt
func Crypt(password string) (string, error) {
	// Hash password with bcrypt
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(encryptPassword), err
}

// VerifyPassword Verify the password is consistent with the hashed password in the database
func VerifyPassword(password, encryptPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
