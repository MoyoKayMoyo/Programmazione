package main

import . "fmt"

func main() {
	var n int
	Print("Insert a number: ")
	Scan(&n)

	if n <= 0 {
		Print("number not positive")
	} else {
		primeTwinNumber(n)
		Println()
	}
}

func isPrime(n int) bool {
	for d := 2; d < n; d++ {
		if n%d == 0 {
			return false	
		}
	}
	return true
}

func primeTwinNumber(limit int) {
	var p int
	Println("twin number inferior to", limit, ":")
	for q := 2; q < limit; q++ {
		if isPrime(q){
			p = q+2
			if isPrime(p) {
				Print("(",q,",",p,")")
			}
		}
	}
}
