package main

import . "fmt"

func main(){
	var a, b, somma int

	Print("Inserisci il primo numero: ")
	Scan(&a)
	Print("Inserisci il sencodo numero: ")
	Scan(&b)

	for i := a; i <= b; i++ {
		if i%2 == 0 {
			somma += i
			Println(i)
		}
		//"%0.17f\n" 17 ciffre dopo la virgola
	}

	Println("La somma dei numeri completi equivale a", somma)
}