package util

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash returns a hash from a string
func Hash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash)
}

// CompareHash checks a hash and a string to see if they're the same
func CompareHash(hash, check string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(check))
	return err == nil
}
