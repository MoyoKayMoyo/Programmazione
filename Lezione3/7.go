package main

import . "fmt"

func main() {
	var v int
	Print("Inserire un voto:")
	Scan(&v)

	if v < 60 {
		if v < 0 {
			Println("Errore")
		} else {
			Println("Insufficiente")
		}
	} else if v < 70 {
		Println("Sufficiente")
	} else if v < 80 {
		Println("Buono")
	} else if v < 90 {
		Println("Distino")
	} else if v <= 100 {
		Println("Ottimo")
	} else if v > 100 {
		Println("Errore")
	}
}
