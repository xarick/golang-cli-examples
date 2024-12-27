package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

// go run .\main.go .\hello.go .\calculate.go calc -a 10 -b 3 -op add
// go run .\main.go .\hello.go .\calculate.go hi -name bobur -age 28 -debug true

// go run .\main.go .\hello.go .\calculate.go hi -help

func main() {
	// Subcommand lar
	hiCmd := flag.NewFlagSet("hi", flag.ExitOnError)
	calcCmd := flag.NewFlagSet("calc", flag.ExitOnError)

	// 1-flaglarni aniqlaymiz
	name := hiCmd.String("name", "Anonim", "Foydalanuvchi ismi")
	age := hiCmd.Int("age", 20, "Yoshingiz")
	debug := hiCmd.Bool("debug", false, "Debug rejim")

	// 2-flaglarni aniqlaymiz
	a := calcCmd.Float64("a", 0, "Birinchi son (float64)")
	b := calcCmd.Float64("b", 0, "Ikkinchi son (float64)")
	op := calcCmd.String("op", "add", "Amal turi: add, sub, mul, div")

	// Barcha flaglarni parse qilamiz
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatal("Subcommand ko'rsatilishi shart: hi yoki calc")
	}

	// Subcommand ni aniqlash
	switch flag.Arg(0) {
	case "calc":
		calcCmd.Parse(flag.Args()[1:])
		result := calculate(*a, *b, *op)
		fmt.Printf("Natija: %.2f %s %.2f = %.2f\n", *a, strings.ToUpper(*op), *b, result)
	case "hi":
		hiCmd.Parse(flag.Args()[1:])
		hello(*name, *age, *debug)
	default:
		log.Fatalf("Noto'g'ri subcommand: %s", flag.Arg(0))
	}
}
