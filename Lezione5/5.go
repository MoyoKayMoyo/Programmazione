package main

import . "fmt"

func main() {
	var n int
	Print("Insert a number: ")
	Scan(&n)

	if n < 0 {
		Print("number not positive")
	} else {
		primeNumber(n)
		Println()
		Println()
	}
}

func isPrimo(n int) bool {
	for d := 2; d < n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func primeNumber(n int) {
	Print("prime number under ", n, ":\n")
	for d := 2; d < n; d++ {
		if isPrimo(d) {
			Print(d, " ")
		}
	}
}
