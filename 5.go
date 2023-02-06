package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main(){
	slt := readNumbers()
	Println("Minumum: ", min(slt))
	Println("Maximum: ", max(slt))
	Printf("Avarage: %0.2f\n", avarage(slt))

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

func min(slt []int) int{
	min := slt[0]
	for _, v := range slt {
		if min > v {
			min = v
		}
	}
	return min	
}

func max(slt[]int) int{
	max := slt[0]
	for _, v := range slt {
		if max < v {
			max = v
		}
	}
	return max
}

func avarage(slt[]int) float64{
	sum := 0
	for _, v := range slt {
		sum += v
	}

	return float64(sum)/float64(len(slt))
}