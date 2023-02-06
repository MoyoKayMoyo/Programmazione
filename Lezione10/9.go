package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"unicode"
)

func main(){
	printHistogram(occ(readText()))
	
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

func printHistogram(histogram map[rune]int){
	keys := make([]string, len(histogram))
	for k := range histogram {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	Println("\nHistogram:")
	for _, s := range keys {
		for k, v := range histogram {
			if s == string(k){
				Print(s, " : ")
				for i := 0; i < v; i++ {
					Print("*")
				}
				Println()
				break
			} 
		}
	}
}