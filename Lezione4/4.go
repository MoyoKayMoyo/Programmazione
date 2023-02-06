package main

import . "fmt"

func main(){
	var num int

	Print("Inserisci un numero: ")
	Scan(&num)

	for i := 0; i <= 10; i++ {
		Println(num,"x", i, "=", num*i)
	}
}
