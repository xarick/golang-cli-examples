package internal

import (
	"regexp"
)

// Email validatsiya qilish uchun regex
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Email manzilini tekshirish funksiyasi
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}
