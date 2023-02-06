package main

import . "fmt"

func main()  {
	var n int
	Print("Insert a number: ")
	Scan(&n)

	printChessboard(n)
}

func printStartStar(lenght int){
	for i := 0; i < lenght; i++ {
		if i%2 == 0{
			Print("*")
		} else {
			Print("+")
		}	
	}
}

func printStartPlus(lenght int){
	for i := 0; i < lenght; i++ {
		if i%2 == 0{
			Print("+")
		} else {
			Print("*")
		}
	}
}

func printChessboard(size int)  {
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			printStartStar(size)
		} else {
			printStartPlus(size)
		}
		Println()
	}
}