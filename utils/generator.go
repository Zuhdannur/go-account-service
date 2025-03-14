package utils

import (
	"math/rand"
	"time"
)

func GenerateAccountNumber(length int) string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	const digits = "0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = digits[rand.Intn(len(digits))]
	}

	return string(result)
}
