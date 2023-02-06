package main

import . "fmt"

func main() {
	var n int
	Print("Inserisci un numero (maggiore di 1): ")
	Scan(&n)

	if n > 1 {

		for r := 0; r < n*2; r++ {

			for c := 0; c < n+(n-1); c++ {

				if c == n-1 {
					Print("* ")
				} else if c < n-1 && r < n-1 {
					if r == 0 || c == r {
						Print("* ")
					} else {
						Print("  ")
					}
				} else if c >= n && r > n {

					if r == 2*n-1 || c == r-1 {
						Print("* ")
					} else {
						Print("  ")
					}

				} else {
					Print("  ")
				}
			}
			Println("  ")
		}

	} else {
		Println("I told u, l'unica cosa che non dovevi fare l'hai fatta")
	}
}
