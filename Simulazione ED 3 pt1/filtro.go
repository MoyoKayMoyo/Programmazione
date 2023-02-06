package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main(){

	// non è stato inserito niente nella riga di codice
	if len(os.Args) == 1 {
		fmt.Println("non è stata insertita nessuna codifica")
	} else {
		// controllo che os.Args abbia una lunghezza pari, ovvero che ci siano solo coppie ch, n
		if len(os.Args[1:])%2 != 0 {
			fmt.Println("coppia di codifiche incompleta")
		} else {
			// estraggo solo i caratteri
			// si trovano nelle posizioni dispari della slice os.Args
			var chars []rune
			for i := 1; i < len(os.Args); i+=2 {
				e := []rune(os.Args[i])
				chars = append(chars, e...)
			}
			// controllo che le stringhe inserite siano effettivamente caratteri
			ch, valid := isCharValid(chars)

			if valid{
				// estraggo i numeri che si trovano nelle posizioni pari della slice os.Args
				var nums []int
				for i := 2; i < len(os.Args); i+=2 {
					n, err := strconv.Atoi(os.Args[i])
					if err != nil {
						fmt.Println(os.Args[i], "non è un numero")
						break
					}

					// controllo che il numero di una coppia non sia minore di 0
					// nel caso fosse 0, la coppia non verrà presa in considerazoine
					if n <= 0 {
						fmt.Println("la lunghezza di ", os.Args[i-1], "invalida, non verrà presa in considerazione")
						break
					}
					nums = append(nums, n)
				}
				// stampo la stringa decodificata
				fmt.Println("Stringa decodificata:")
				for i := 0; i < len(nums); i++ {
					for k := 0; k < nums[i]; k++ {
						fmt.Print(ch[i])
					}
				}
				fmt.Println()
			}
		}
	}
}

// funzione che il carattere sia valido
func isCharValid(s []rune) (ch []string, valid bool){
	for _, v := range s {
		if !unicode.IsLetter(v){
			fmt.Println("Elemento", string(v), "deve essere un carattere")
			return ch, false
		} else {
			ch = append(ch, string(v))
		}
	}
	return ch, true
}
