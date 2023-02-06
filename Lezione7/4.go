package main

import (
	. "fmt"
	"strings"
)

func main() {

	for {
		var s string
		Print("Insert a string: ")
		Scanln(&s)
		if s == "" {
			Println()
			break
		}
		Println("String in uppercase:", strings.ToUpper(s))
		Println()
	}
	Println("Bye!")
}
