package main

import . "fmt"

func main() {
	var num int
	Print("Inserire un numero: ")
	Scan(&num)

	Print("divisori di ", num, ": ")
	for d:=1; d<num;d++{
		if num%d == 0 {
			Print(d, " ")
		}
	}
	Println()
}
