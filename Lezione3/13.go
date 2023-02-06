package main

import . "fmt"

func main(){
	var m, q, px, py float64

	Print("Inserire m:")
	Scan(&m)
	Print("Inserire q:")
	Scan(&q)

	Print("\nInserire px:")
	Scan(&px)
	Print("Inserire py:")
	Scan(&py)

	y := m * px + q
	Println()

	if y == py {
		Println("Il punto fa parte della retta")
	} else if y < py {
		Println("Il punto si trova al di sotto della retta")
	} else {
		Println("Il punto si trova al di sopra della retta")
	}

}