package main

import . "fmt"

func main(){
	var n, num int
	var numbers[]int
	Printf("Insert the number of integers: ")
	Scan(&n)

	Printf("Insert %d numbers: \n", n)
	for i := 0; i < n; i++ {
		Scan(&num)
		numbers = append(numbers, num)
	}

	Println("numbers in reverse order")
	for i := len(numbers)-1; i >= 0; i-- {
		Printf("%d ", numbers[i])
	}
	Println()
}