package main

import (
	"os"
	"xarick/cli/cli-app/internal"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "Password Generator",
		Usage:    "Turli xil parollar generatsiya qiladi",
		Commands: internal.GetCommands(),
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

// go run main.go pwd --size 12 --type number
// go run main.go pwd --size 15 --type letter
// go run main.go pwd -s 20 -t all

// go run main.go validate --email test123@example.com
