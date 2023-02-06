package main

import . "fmt"

func main() {
	var a1, a2 float64
	Print("Inserire angolo 1: ")
	Scan(&a1)

	Print("Inserire angolo 2: ")
	Scan(&a2)

	if a1+a2 < 180 {
		a3 := 180 - (a1 + a2)
		Print("Ampiezza terzo angolo: ", a3, "°\n")
	} else {
		Println("I due angoli non appartengono a un triangolo")
	}
}
