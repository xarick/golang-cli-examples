package internal

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func GetCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "pwd",
			Usage: "Parol generatsiya qiladi",
			Flags: []cli.Flag{
				&cli.IntFlag{Name: "size", Aliases: []string{"s"}, Usage: "Parol uzunligi", Required: true},
				&cli.StringFlag{Name: "type", Aliases: []string{"t"}, Usage: "Parol turi (number, letter, all)", Required: true},
			},
			Action: func(c *cli.Context) error {
				size := c.Int("size")
				pType := c.String("type")

				password, err := GeneratePassword(size, pType)
				if err != nil {
					return err
				}

				fmt.Printf("Yangi parol: %s\n", password)
				return nil
			},
		},
		{
			Name:  "validate",
			Usage: "Email manzilini tekshiradi",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "email", Aliases: []string{"e"}, Usage: "Tekshiriladigan email manzili", Required: true},
			},
			Action: func(c *cli.Context) error {
				email := c.String("email")

				if ValidateEmail(email) {
					fmt.Println("To'g'ri email")
				} else {
					fmt.Println("Noto'g'ri email")
				}
				return nil
			},
		},
	}
}
