package main

import (
	. "fmt"
	"strconv"
	"os"
)

func main(){

	numbers := readNumbers()
	mult := 1 //neutral number for the multiplication
	for _, v := range numbers {
		mult *= v
	}
	Println("The result of the multiplication is", mult)

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