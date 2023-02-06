package main

import . "fmt"

func main() {
	var num, somma, a, min, max, pos, neg, nulli int

	Print("Inserisci un numero: ")
	Scan(&num)

	Println("Inserisci", num, " numeri")
	for i := 0; i < num; i++ {
		Scan(&a)
		somma += a

		if i == 0 {
			min = a
			max = a
		}

		if a < min {
			min = a
		}

		if  a > max {
			max = a
		}

		if a > 0 {
			pos++
		}

		if a < 0 {
			neg++
		}

		if a == 0 {
			nulli++
		}
	}

	Println("somma = ", somma)
	Println("valore minimo = ", min)
	Println("valore massimo = ", max)
	Println("interi positivi = ", pos)
	Println("interi negativi = ", neg)
	Println("interi nulli = ", nulli)


}