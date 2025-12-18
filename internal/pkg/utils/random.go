package utils

import (
	"crypto/rand"
	"math/big"
)

// RandomNumber generates secure random number between min & max
func RandomNumber(min, max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()) + min, nil
}

// OTP generates numeric OTP of given length
func OTP(length int) (string, error) {
	const digits = "0123456789"

	result := make([]byte, length)
	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		result[i] = digits[n.Int64()]
	}

	return string(result), nil
}
