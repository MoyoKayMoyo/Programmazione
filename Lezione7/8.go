package main

import (
	. "fmt"
	"strconv"
)

type Fraction struct {
	numerator int
	denominator int
}

func main(){
	var n, d int
	Print("Insert two numbers: ")
	Scan(&n, &d)

	Println("Fraction:", String(*newFraction(n,d)))
}

func newFraction(numerator, denominator int) *Fraction {
	var f *Fraction
	f = &Fraction{numerator, denominator}
	
	return f
}

func String(f Fraction) string {
	n, d := strconv.Itoa(f.numerator), strconv.Itoa(f.denominator)
	s := n + "/" + d

	return s
}