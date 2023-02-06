package main

import . "fmt"
import math "math"

func main (){
	var x, y int
	Print("Inserisci x: ")
	Scan(&x)
	Print("Inserisci y: ")
	Scan(&y)
	Println()

	if x > y {
		Println("Maggiore:", x)
		Println("Minore:", y)
		Println("Differenza:", x-y)
	}

	if x < y {
		Println("Maggiore:", y)
		Println("Minore:", x)
		Println("Differenza:", y-x)
	}

	Println("Somma:", y)
	Println("Prodotto:", x*y)
	Println("Divisione:",float64(x)/float64(y))
	Println("Valore medio:", float64((x+y))/float64(2))

	Px := x
	for i := y; i > 1; i-- {
		Px *= x
	}
	Println("Potenza (for) : ", Px)
	Println("Potenza (math.pow): ", math.Pow(float64(x), float64(y)))
}