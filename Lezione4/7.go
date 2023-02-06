package main

import . "fmt"

func main() {
	var num, somma int
	Print("Inserisci una sequenza di numeri: ")
	
	i := 0
	for ;;i++ {
		Scan(&num)
		if num <= 0{
			break
		}
		somma += num
	}

	Println("media = ", float64(somma)/float64(i))
}