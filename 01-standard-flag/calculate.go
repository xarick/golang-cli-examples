package main

import "log"

func calculate(a, b float64, op string) float64 {
	switch op {
	case "add":
		return a + b
	case "sub":
		return a - b
	case "mul":
		return a * b
	case "div":
		if b == 0 {
			log.Fatal("Xatolik: Nolga bo'lish mumkin emas!")
		}
		return a / b
	default:
		log.Fatalf("Noto'g'ri amal: %s", op)
		return 0
	}
}
