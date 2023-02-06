package main

import (
	. "fmt"
	rand "math/rand"
	"time"
)

func main(){

	/* inizializzazione del generatore di numeri casuali */
	rand.Seed(int64(time.Now().Nanosecond())) 
	/* generazione di un numero casuale compreso nell'intervallo 
	che va da 0 a 99 (estremi inclusi) */
	var numeroGenerato int = rand.Intn(100)+1
	var a int

	for i:= 0;;i++{
		Print("Tentativo n°", i+1, ": ")
		Scan(&a)

		if a == numeroGenerato {
			Println("hai indovinato in", i , "tentativi!")
			return
		}

		if a > numeroGenerato{
			Println("Il numero inserito è più basso! RIPROVA")
		} else {
			Println("Il numero inserito è più alto! RIPROVA")
		}
	}

}