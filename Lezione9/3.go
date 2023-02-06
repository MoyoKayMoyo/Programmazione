package main

import (
	. "fmt"
	"os"
)

func main(){

	if !isBrackets(os.Args[1]){
		Println("the string inserted does not have only round brackets")
	} else {
		if isBalanced(os.Args[1]){
			Println("brackets balanced")
			subBrackets(os.Args[1])
		} else {
			Println("brackets not balanced")
			subBrackets(os.Args[1])
		}

	}
}

func isBalanced(b string) bool{
	var open, close int
	for _, br := range b {
		if string(br) == "(" {
			open++
		} else if string(br) == ")" && open > 0 && open != close{
			close++
		}
	}
	if open == close {
		return true
	}

	return false
}

func isBrackets(b string) bool{
	for _, br := range b { 
		if string(br) != "(" && string(br) != ")"{
			return false
		}
	}
	return true
}

func subBrackets(b string){
	n := 1
	for size := 0; size < len(b); size++ {
		for i := size + 2; i <= len(b); i++ {
			if len(b[size:i])%2 == 0 && string(b[size]) != ")" && isBalanced(string(b[size:i]) ){
				Println(n, ":", string(b[size:i]))
				n++
			}
		}
	}

	if n == 1 {
		Println("there are not substring")
	}
}