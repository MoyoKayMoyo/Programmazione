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
	Reduce(f)
	return f
}

func String(f Fraction) string {
	n, d := strconv.Itoa(f.numerator), strconv.Itoa(f.denominator)
	s := n + "/" + d

	return s
}

func Reduce (f *Fraction) {
	var lenght int
	if f.numerator > f.denominator {
		lenght = f.denominator
	} else {
		lenght = f.numerator
	}

	for i := lenght; i >= 2; i-- {
		if f.numerator%i == 0 && f.denominator%i == 0{ 
			f.numerator, f.denominator = f.numerator/i, f.denominator/i
		}
	}
}