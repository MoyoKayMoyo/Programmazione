package main

import . "fmt"

func main() {
	var s string
	Print("Insert a string: ")
	Scan(&s)

	var lower, upper, consonant, vocal int
	for _, r := range s {

		if isValid(r){
			if isUppercase(r) {
				upper++
			} else if isLower(r) {
				lower++
			}
			
			if isVocal(r){
				vocal++
			} else {
				consonant++
			}
		}
	}

	Println("Uppercase:\t", upper)
	Println("Lowercase:\t", lower)
	Println("Vocal:	\t", vocal)
	Println("Consonant:\t", consonant)
}

func isValid(r rune) bool {
	if isUppercase(r) || isLower(r) {	
		return true
	} 
	return false
}

func isUppercase(r rune) bool {
	if r >= 'A' && r <= 'Z' {	
		return true
	} 
	return false
}

func isLower(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	} 
	return false
}

func isVocal(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}

	

	return false
}
