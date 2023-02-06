package main

import . "fmt"
import math "math"

func main() {
	var y1, y2 float64 // y1 = year person 1, y2 = year person 2

	Print("Inserire l'età della prima persona:")
	Scan(&y1)
	Print("Inserire l'età della prima persona:")
	Scan(&y2)

	Println("Somma età:", y1+y2)
	Println("Media età:", (y1+y2)/2)
	Println("La media delle età arrontondata per difetto all'intero inferiore:", math.Floor((y1 + y2)/2))
	Println("La media delle età arrontondata per eccesso all'intero superiore:", math.Ceil((y1+y2)/2))
	Println("Somma età tra 10 anni:", y1+y2+20)
	Println("Media età tra 10 anni:", (y1+y2+20)/2)
}
