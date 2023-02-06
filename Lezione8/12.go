package main

import (
	. "fmt"
	"os"
	"sort"
	"strconv"
)

func main(){
	if !readNumbers(){
		Println("the given element is not valid")
	} else {
		numero := os.Args[1]
		var sl[]string
		sl = append(sl, numero)
		num := ""
		number := make([]string, 0)

		for _, v := range numero {
			number = append(number, string(v))
		}
	
		for i := 0; i < len(number); i++ {
			for j := 0; j < len(number); j++ {
				if j != i {
					num += number[j]
				}
			}
			if len(num) != 0 {
				sl = append(sl, num)
				num = ""
			}
		}
		for i := 0; i < len(number); i++ {
			for j := 0; j < len(number); j++ {
				if j != i && j!= i+1 && i+1 < len(number){
					num += number[j]
				}
			}
			if len(num) != 0 {
				sl = append(sl, num)
				num = ""
			}
		}
		for i := 0; i < len(number); i++ {
			for j := 0; j < len(number); j++ {
				if j != i && j!= i+1 && i+1 < len(number) && j!= i+2 && i+2 < len(number){
					num += number[j]
				}
			}
			if len(num) != 0 {
				sl = append(sl, num)
				num = ""
			}
		}
		lista := covertNumbers(sl)
		for _, v := range lista {
			if isPrime(v){
				Println(v)
			}
		}
	}
}

func isPrime(n int) bool{
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func readNumbers() bool{
	if len(os.Args[1]) >= 4 {
		_, err := strconv.Atoi(os.Args[1])
		if err == nil{
			return true
		}
	}
	return false
}

func covertNumbers(sl []string) (numbers[]int) {
	for _, v := range sl {
		num, err := strconv.Atoi(v)
		if err == nil{
			numbers = append(numbers, num)
		}
	}
	sort.Ints(numbers)
	return numbers
}