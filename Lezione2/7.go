package main

import . "fmt"
import "math"

const pi_greco = math.Pi

func main() {
	var r float64 //r = raggio

	Print("Inserire raggio del cerchio:")
	Scan(&r)

	Println("Perimetro cerchio:", 2*r*pi_greco)
	Println("Area cerchio:", r*r*pi_greco)
}
