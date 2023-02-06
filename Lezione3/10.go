package main

import . "fmt"

func main() {
	var a, b int
	Print("Inserire un numero a: ")
	Scan(&a)

	Print("Inserire un numero b: ")
	Scan(&b)

	if b == 0 {
		Println("Impossibile")
	} else {
		Println("Quoziente:", float64(a)/float64(b))
	}
}
