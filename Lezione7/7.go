package main

import (
	. "fmt"
	"strconv"
	"unicode"
)

func main() {
	for {
		var s string
		Print("Pwease insert a string: ")
		Scanln(&s)
		if s == "" {
			Println()
			break
		}
		sum, _ := hiddenNumber(s)
		if sum != 0{
			Println("Double of hidden numbre:", sum*2, "(", sum, "* 2",")")
		} else {
			Println("no hidden numbers")
		}
		Println()
	}

}

func hiddenNumber(s string) (int, error){
	var sum string
	for _, v := range s {
		if unicode.IsDigit(v) {
			sum += string(v)
		}
	}
	if sum == "" {
		sum = "0"
	}
	i, err := strconv.Atoi(sum)
	
	return i, err
}