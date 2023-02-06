package main

import . "fmt"

func main() {
	var b, h float64 // b = base, a = altezza

	Print("Inserire base:")
	Scan(&b)
	Print("Inserire altezza:")
	Scan(&h)

	Println("Perimetro rettangolo:", 2*b+2*h)
	Println("Area rettangolo:", b*h)
}
