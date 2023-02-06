package main

import . "fmt"

func main()  {
	var n int
	Print("Inserisci un numero: ")
	Scan(&n)

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if r%2 == 0 {
				Print("* ")
			} else {
				Print("+ ")
			}
		}
		Println()
	}
}