package main

import . "fmt"

func main()  {
	var n int
	var s string
	Print("Insert a string: ")
	Scan(&s)

	Print("Insert a number: ")
	Scan(&n)

	for i := 0; i < n; i++ {
		Print(s)
		if i!=n-1{
			Print("-")
		}
	}
	Println()
}