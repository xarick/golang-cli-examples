package internal

import (
	"errors"
	"math/rand"
	"time"
)

// Parol generatsiya qilish uchun belgilar to‘plami
const (
	numbers    = "0123456789"
	letters    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	allSymbols = "!@#$%^&*()-_=+[]{}<>?,."
)

// Parol generatsiya qilish funksiyasi
func GeneratePassword(size int, pType string) (string, error) {
	var charset string

	switch pType {
	case "number":
		charset = numbers
	case "letter":
		charset = letters
	case "all":
		charset = numbers + letters + allSymbols
	default:
		return "", errors.New("noto'g'ri parol turi, faqat 'number', 'letter' yoki 'all' bo'lishi mumkin")
	}

	// eski usul
	// rand.Seed(time.Now().UnixNano())

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	password := make([]byte, size)
	for i := range password {
		password[i] = charset[r.Intn(len(charset))]
	}

	return string(password), nil
}
