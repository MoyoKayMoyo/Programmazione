package main

import . "fmt"

func main() {
	var num int
	Print("Inserire un numero: ")
	Scan(&num)

	if num%3 == 0 {
		Println("Fizz")
	}
	if num%5 == 0 {
		Println("Buzz")
	}
	Println()
}
