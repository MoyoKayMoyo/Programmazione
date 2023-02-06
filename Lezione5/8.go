package main

import . "fmt"

func main() {
	var s string
	Print("Insert a string: ")
	Scan(&s)

	Println("Text in uppercase")
	for _, r := range s {
		Print(string(uppercase(r)))		
	}

	Println()

	Println("Text in lowercase")
	for _, r := range s {
		Print(string(lowercase(r)))		
	}
	Println()

}

func uppercase(r rune) rune{
	if r >= 'a' && r <= 'z' {
		return r + 'A' - 'a'
	}
	return r
}

func lowercase(r rune) rune{
	if r >= 'A' && r <= 'Z' {
		return r + 'a' - 'A'
	} 
 
	return r
}