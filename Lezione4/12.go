package main

import . "fmt"

func main() {
	var num, max int	

	Println("Inserisci una sequenza di numeri: ")
	for {
		Scan(&num)
		if num <= -1 {
			break
		}
		if num > max{
			Println("crescente")
		} else if num < max {
			Println("decrescente")
		} else {
			Println("uguale")
		}
		max = num
		Println()
	}
}	