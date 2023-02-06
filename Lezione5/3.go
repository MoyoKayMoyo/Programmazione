package main

import . "fmt"

func main() {
	var n int
	Print("Inserisci un numero: ")
	Scan(&n)

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			switch r % 3 {
			case 0:
				Print("* ")
			case 1:
				Print("+ ")
			case 2:
				Print("o ")
			}
		}
		Println()
	}
}
