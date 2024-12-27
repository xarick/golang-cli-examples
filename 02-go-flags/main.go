package main

import (
	"fmt"
	"os"
	"xarick/cli/go-flags/cli"

	"github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		List  cli.ListCommand     `command:"list" description:"List all todos"`
		Show  cli.ShowCommand     `command:"show" description:"Show todo by ID"`
		Posts cli.ListPostCommand `command:"posts" description:"List all posts"`
	}

	parser := flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		fmt.Printf("Error parsing arguments: %v\n", err)
		os.Exit(1)
	}

	// _, err := flags.Parse(&opts)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}

// go run main.go --help
