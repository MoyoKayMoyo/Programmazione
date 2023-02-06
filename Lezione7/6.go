package main

import . "fmt"

func main() {
	var last string
	for {
		var s string
		Print("Pwease insert a surname: ")
		Scanln(&s)
		if s == "" {
			Println()
			break
		}
		if last < s{
			last = s
		}

	}
	Println("Last Surname: ", last)
}