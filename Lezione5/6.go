package main

import . "fmt"

func main() {
	var n int
	Print("Insert a number: ")
	Scan(&n)

	if n < 0 {
		Print("number not positive")
	} else {
		copiuosNumber(n)
		Println()
	}
}

func sumDiversors(n int) int {
	var sum int
	for d := 1; d < n; d++ {
		if n%d == 0 {
			sum += d
		}
	}
	return sum
}

func isCopious(n int) bool {
	if sumDiversors(n) > n {
		return true
	}
	return false
}

func copiuosNumber(limit int) {
	Print("copious numbers: ")
	for i := 1; i < limit; i++ {
		if isCopious(i) {
			Print(i, " ")
		}
	}
}
