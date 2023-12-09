package util

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// Crypt Encrypt the password using crypto/bcrypt
func Crypt(password string) (string, error) {
	if password == "" {
		return "", errors.New("password为空字符串")
	}
	// Hash password with bcrypt
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(encryptPassword), err
}

// VerifyPassword Verify the password is consistent with the hashed password in the database
func VerifyPassword(password, encryptPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
}
