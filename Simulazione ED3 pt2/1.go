package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main(){
	fmt.Println("Inserire una stringa: ")
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		text = scanner.Text()
		break
	}

	valid, s := isValid(text)
	if valid {
		if isPalindroma(s){
			fmt.Println("la frase è palindroma")
		} else {
			fmt.Println("la frase non è palindroma")
		}
		
	}

}

func isValid(text string)(bool, string){
	alfabeto := "abcdefghilmnopqrstuvzàèìòù"
	text = strings.ToLower(text)
	var s string
	for _, v := range text {
		if unicode.IsLetter(v){
			if strings.Contains(alfabeto, string(v)){
				s += string(v)
			} else {
				fmt.Println(string(v))
				fmt.Println("La stringa immessa non contiene solo caratteri italiani")
				return false, s
			}
		}
	}
	return true, s
}

func isPalindroma(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}