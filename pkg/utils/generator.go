package utils

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateCode() string {

	length := 6

	code := make([]byte, length)

	for i := range code {

		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))

		code[i] = charset[n.Int64()]
	}

	return string(code)
}
