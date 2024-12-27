package main

import (
	"fmt"
)

func hello(name string, age int, debug bool) {
	fmt.Printf("Salom, %s!\n", name)
	fmt.Printf("Yoshingiz: %d\n", age)

	if debug {
		fmt.Println("Debug rejimi yoqildi!")
	} else {
		fmt.Println("Debug rejimi o'chirilgan.")
	}
}
