package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){

	for i, v := range os.Args[1:] {
		fmt.Print(trasformaParola(v, i), " ")
	}
	fmt.Println()
}

func trasformaParola(parola string, posizione int) (parolaTrasformata string){
	for i := 0; i < len(parola); i++ {
		if posizione%2 == 0{
			if i%2 == 0{
				parolaTrasformata += strings.ToUpper(string(parola[i]))
			} else {
				parolaTrasformata += strings.ToLower(string(parola[i]))
			}
		} else {
			if i%2 == 1{
				parolaTrasformata += strings.ToUpper(string(parola[i]))
			} else {
				parolaTrasformata += strings.ToLower(string(parola[i]))
			}
		}
	}
	return
}