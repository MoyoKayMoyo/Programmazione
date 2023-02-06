package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main(){

	num := occurences(readNumbers())
	Println("Result:", num)

}

func readNumbers()(numbers[]int) {
	for _, v := range os.Args [1:]{
		num, err := strconv.Atoi(v)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return
}

func occurences(numbers[]int) int {
	sum := 0
	var double []int
	for i, v := range numbers {
		for j := i+1; j < len(numbers); j++ {
			if v == numbers[j]{
				sum += numbers[j]
				double = append(double, numbers[j])
			}
		}
	}
	Println(double)
	return sum
}