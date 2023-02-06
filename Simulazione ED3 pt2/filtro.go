package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	if len(os.Args) == 3 {
		a, err := strconv.Atoi(os.Args[1])

		if err == nil {
			b, err := strconv.Atoi(os.Args[2])

			if err == nil {

				if a >= b {
					fmt.Println("NaN")
				} else {
					somma := 0
					counter := 0
					for i := a; i <= b; i++ {
						if i%2 == 1 {
							somma += i
							counter++
						}
					}
					fmt.Println(somma/counter)
				}

			} else {
				fmt.Println("il secondo argomento non è un numero")
				os.Exit(1)
			}

		} else {
			fmt.Println("il primo argomento non è un numero")
			os.Exit(1)
		}

	} else {
		fmt.Println("il numero di argomenti inseriti non è corretto")
	}
}
