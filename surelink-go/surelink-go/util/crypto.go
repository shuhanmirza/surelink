package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func CheckHash(text, hashString string) bool {
	hashText := hex.EncodeToString(GetHash([]byte(text)))

	return hashText == hashString
}
