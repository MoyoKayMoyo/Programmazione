package main

import (
	. "fmt"
	"bufio"
	"os"
	"unicode"
)

func main(){
	words_count := countRipetition(divideWords(readText()))
	Println("\nDistint words:", len(words_count))
	for k, v := range words_count {
		Println(k, "\t:", v)
	}
}

func readText() (txt string){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt += scanner.Text()
	}

	return
}

func divideWords(s string) (words []string){
	word := ""
	isWord := false
	for i, v := range s {
		if unicode.IsLetter(v){
			word += string(v)
		} else if len(word) != 0{
			isWord = true
		}

		if isWord  && len(word) != 0 || !isWord && i == len(s)-1{
			//Println(i, word)
			words = append(words, word)
			isWord = false
			word = ""
		}
	}

	return
}

func countRipetition(sl []string) map[string]int{
	occ := make(map[string]int)
	for _, v := range sl {
		_, ok := occ[v]
		if !ok {
			occ[v] = 1
		} else {
			occ[v]++
		}
	}
	return occ
}