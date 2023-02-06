package main

import (
	. "fmt"
	"math"
)

func main()  {
	var n int
	Print("Pwease insert a number:")
	Scan(&n)
	isPitaghoreanTriple(n)
	
}

func isTriple(a, b, c int) bool {
    quad := float64(2)
    if math.Pow(float64(c), quad) ==  math.Pow(float64(a), quad) + math.Pow(float64(b), quad) {
        return true
    }
    return false
}

func isPitaghoreanTriple(n int) {
    for i := 1; i < n; i++ {
        for j := 1; j < n; j++ {
            for k := 1; k < n; k++ {
                if isTriple(i, j, k){
                    Printf("(%d, %d, %d)\n", i,j,k)
                }
            }
        }
    }
}