package main

import (
	. "fmt"
	"os"
)

func main() {

	if len(os.Args[1:]) == 1 {
		text := []rune(os.Args[1])
		for size := 0; size < len(text); size++ {
			for i := size + 2; i <= len(text); i++ {
				if isPalindrome(text[size:i]) {
					Println("word:", string(text[size:i]), "size:", len(text[size:i]))
				}
			}
		}
	} else {
		Println("More than one string inserted uwu")
	}

}

func isPalindrome(text []rune) bool {
	// pal := ""
	// for _, v := range text {
	// 	pal = string(v) + pal

	// }
	// if pal == text {
	// 	return true
	// }
	// return false

	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-i-1] {
			return false
		}
	}

	return true
}
