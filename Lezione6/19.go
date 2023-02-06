package main

import . "fmt"

func main(){
	var s1, s2 string
	Print("Insert the first string:")
	Scan(&s1)
	Print("Insert the second string:")
	Scan(&s2)

	Println("Alternated string: ", alternateString(s1, s2))
}

func alternateString (s1, s2 string) string {
	var s3 string
	var lenght int

	if len(s1) > len(s2) {
		lenght = len(s1)
	} else {
		lenght = len(s2)
	}


	for i := 0; i < lenght; i++ {
		if i < len(s1) {
			s3 += string(s1[i])
		} else {
			s3 += "-"
		}

		if i < len(s2) {
			s3 += string(s2[i])
		} else {
			s3 += "-"
		}
	}
	return s3
}