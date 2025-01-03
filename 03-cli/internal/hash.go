package internal

import (
	"crypto/sha256"
	"encoding/hex"
)

// Matnni hash qilish funksiyasi
func GenerateHash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
