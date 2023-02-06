package main

import . "fmt"

func main(){
	var s string
	Print("Insert a string: ")
	Scan(&s)

	for i := 0; i < len(s); i++ {
		Print(string(s[i])," ")
	}
	Println()
}