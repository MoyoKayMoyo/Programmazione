package main

import . "fmt"

func main() {
	var km float64
	mi := 0.62137

	Print("Inserire Km:")
	Scan(&km)
	Println(km, "km sono", km*mi, "miglia")
}
