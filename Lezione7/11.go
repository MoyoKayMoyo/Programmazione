package main

import (
	. "fmt"
	"math"
	"strings"
)

func main(){
	var n string
	var b int
	Print("insert a number in base 2 to 16: ")
	Scan(&n)
	Print("insert the base of the number: ")
    Scan(&b)

	s, e := xTo10(n, b)
	if e {
		Println(n, "in base", b, "is", s, "in base 10")
	} else {
		Println("the base", b, "is not valid")
	}
}

func xTo10(n string, b int) (int, bool) {
	if b < 2 || b > 16 {
        return 0, false
    }

	s := "0123456789ABCDEF"
	var num float64
	for i := 0; i < len(n); i++ {
		cb := len(n) - 1 - i
		idx := strings.Index(s, string(n[i]))
		num += float64(idx) * math.Pow(float64(b), float64(cb))
	}

	return int(num), true
}