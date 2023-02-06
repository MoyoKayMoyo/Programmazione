package main

import . "fmt"

func main(){
	var num int

	Print("Inserisci un numero: ")
	Scan(&num)

	for i := 1; i <= num; i++ {

		if i%3 == 0 && i%5 == 0{
			Print("FizzBuzz ")
		}else if i%3 == 0{
			Print("Fizz ")
		}else if i%5 == 0 {
			Print("Buzz ")
		} else {
			Print(i, " ")
		}
	}
	Println()
}