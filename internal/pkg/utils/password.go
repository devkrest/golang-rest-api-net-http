package utils

import "golang.org/x/crypto/bcrypt"

const defaultCost = bcrypt.DefaultCost

// HashPassword hashes plain password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		defaultCost,
	)
	return string(hash), err
}

// ComparePassword compares hash & plain password
func ComparePassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
