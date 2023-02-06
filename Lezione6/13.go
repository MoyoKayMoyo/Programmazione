package main

import . "fmt"

func main(){
	var n int
	Print("Insert a number: ")
	Scan(&n)

	if !littleTable(n){
		Println("Program terminated...")
	}
}

func littleTable(n int) bool {
	if n > 0 && n < 10 {
		Print("Little table of ", n,": ")
		for i := 0; i < 10; i++ {
			Print(i*n, " ")
		}
		Println()
		return true
	} 
	return false
}
