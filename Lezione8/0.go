package main

import (
	. "fmt"
	"os"
)

func main(){
	var alphabet []string
	for i := 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, string(i))
	}

	occ := occ(os.Args[1])
	for i := 0; i < len(occ); i++{
		Print(alphabet[i], ": ")
		Println(occ[i])
	}
}

func occ (s string) (alphabet[26] int){
	for _, v := range s {
		if v <= 'z' && v >= 'a'{
			alphabet[v-'a']++
		}
	}
	return
}
