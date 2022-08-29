package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstwxyzABCDEFGHIJKLMNOPQRSTUVXYZ0123456789="

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomBool() bool {
	if rand.Intn(2) == 0 {
		return false
	}
	return true
}

func RandomStringAlphabet(n int) string {
	var stringBuilder strings.Builder

	lenAlpha := len(alphabet)
	for i := 0; i < n; i++ {
		char := alphabet[rand.Intn(lenAlpha)]
		stringBuilder.WriteByte(char)
	}

	return stringBuilder.String()
}
