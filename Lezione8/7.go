package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main(){

	Println("somma prodotti pari", calculate(readNumbers()))

}

func readNumbers() []int{
	var numbers[]int
	for _, v := range os.Args [1:]{
		num, err := strconv.Atoi(v)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func calculate(sl[]int) int{
	var sum int
	for i, v := range sl {
		for j := sl[i]; j < len(sl); j++ {
			prod := v*sl[j]
			if prod%2 == 0 {
				sum += prod
			}
		}
	}
	return sum
}