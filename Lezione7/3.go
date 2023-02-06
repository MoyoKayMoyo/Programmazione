package main

import . "fmt"

func main()  {
	var sumString string

	for {
		var s string
		Print("Insert a string: ")
		Scanln(&s)
		if s == ""{
			break
		}
		sumString += s + " "

	}
	Println(sumString)
	
}