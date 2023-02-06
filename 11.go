package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main(){
	sum := 0
	numbers := readNumbers()

	for i := 0; i < len(numbers); i++ {
		if occ(numbers, numbers[i]) == 1{
			sum += numbers[i]
		}
	}

	Println("The sum of the unique numbers is", sum)
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

func occ(numbers[]int, n int) (occ int){
	for _, v := range numbers {
		if n == v{
			occ++
		}
	}
	return
}