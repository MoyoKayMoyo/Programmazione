package main

import . "fmt"
import uni "unicode"

func main(){
	var s string
	Print("Insert a string: ")
	Scan(&s)

	var upperCon, upperVocal, lowerCon, lowerVocal int

	for _, r := range s {
		// controlla se è una lettera
		if uni.IsLetter(r){

			// lettera è maiuscola
			if uni.IsUpper(r) {

				// lettera maiuscola è una vocale
				if isVocal(r) {
					upperVocal++
				} else {
					upperCon++
				}

			// lettera minuscola
			} else {
				// lettera minuscola è una vocale
				if isVocal(r){
					lowerVocal++
				} else {
					lowerCon++
				}
			}
		}	
	}

	Println("Upper vocal:	\t", upperVocal)
	Println("Upper consonant:\t", upperCon)
	Println("Lower vocal:	\t", lowerVocal)
	Println("Lower consonant:\t", lowerCon)
}

func isVocal(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
