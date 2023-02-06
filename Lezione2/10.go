package main

import . "fmt"
import math "math"

func main() {
	var n, l int
	Print("Inserisci il numero di lati del poligono:")
	Scan(&n)
	Print("Inserisci la lunghezza dei lati del poligono:")
	Scan(&l)

	Println("Area calcolata:", ((float64(n) * math.Pow(float64(l), 2)) / (4 * math.Tan(math.Pi/float64(n)))))
}
