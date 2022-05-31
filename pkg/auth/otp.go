package auth

import (
	"crypto/rand"
)

const (
	OtpLength = 4 // Длина ОТП
)

func GenerateRandomNumbers(n int) string {
	const letters = "0123456789"

	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return ""
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes)
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateNewOTP() string {
	return GenerateRandomNumbers(OtpLength)
}
