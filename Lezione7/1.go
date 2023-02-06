package main

import . "fmt"

func main(){
	var n int
	s := "*"
	Print("Insert a number: ")
	Scan(&n)

	if n <= 0 {
		Println("Size not sufficient")
	} else {
		for i := 0; i < n; i++ {
			Println(s)
			s+="*"
		}
	}

}
