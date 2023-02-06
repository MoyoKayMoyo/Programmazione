package main

import . "fmt"

func main () {
	var phrase string
	Print("Insert string: ")
	Scan(&phrase)

	if isPalindrome(phrase) {
		Println("Is palindrom")
	} else {
		Println("Not palindrome")
	}

}

func isPalindrome(phrase string) bool{
	var inverse string
	for _, s := range phrase {
		inverse = string(s) + inverse
	}

	if phrase == inverse {
		return true
	}

	return false
}