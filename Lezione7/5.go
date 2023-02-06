package main

import (
	. "fmt"
	"strconv"
)

func main() {
	var s string
	var sum int
	for {
		Print("Pwease insert an integer: ")
		Scanln(&s)
		i, err := strconv.Atoi(s)
		if err != nil {
			break
		}
		sum += i

	}
	Println("Sum:", sum)
}
