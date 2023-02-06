package main

import (
	"bufio"
	. "fmt"
	"os"
	"unicode"
)

func main(){
	histogram := occ(readText())
	Println("\nHistogram:")
	for k, v := range histogram {
		Print(string(k), " : ")
		for i := 0; i < v; i++ {
			Print("*")
		}
		Println()
	}
}

func readText() (txt string){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt += scanner.Text()
	}
	return
}

func occ(s string)(map[rune]int){
	occ := make(map[rune]int)
	for _, v := range s {
		if unicode.IsLetter(v){
			_, ok := occ[v]
			if !ok {
				occ[v] = 1
			} else {
				occ[v]++
			}
		}
	}
	return occ
}
