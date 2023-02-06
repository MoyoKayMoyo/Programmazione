package main

import . "fmt"

func main() {
	var num int
	Print("Inserisci un numero:")
	Scan(&num)

	if num%10 == 0 {
		Println(num, "è un multiplo di 10")
	} else {
		Println(num, "non è un multiplo di 10")
	}
}