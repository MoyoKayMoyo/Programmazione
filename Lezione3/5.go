package main

import . "fmt"

func main() {
	var num int
	Print("Inserisci un numero intero: ")
	Scan(&num)

	if num >= 0 {
		if num == 0 {
			Println(num)
		}
		Print("+", num, "\n")
	} else {
		Print(num, "\n")
	}
}
