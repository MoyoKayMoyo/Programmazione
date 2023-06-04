package main

import (
	"fmt"
	"unicode"
)

/*	Stampare il numero di parole contenute in una frase	*/

func main(){

	s := "ciao come stai, tutto bene?"	// frase
	fmt.Println(contaParole(s))		//stampa il numero di parole

}

func contaParole(s string) (n int){ 		//funzione contaParole che prende una stringa e restituisce il numero di parole
	var isParola bool			//booleano che mi serve per capire se il carattere corrente è una lettera
	for _, v := range s {			//itero sulla frase
		if unicode.IsLetter(v){		//controllo se la runa corrente è una lettera (e non un caattere speciale o altro)
			isParola = true		//se il carattere corrente è una lettera allora isParola è true
		} else if isParola {		//se invece il carattere corrente non è una lettera e isParola è true
			n++			//a parola è terminata e quindi viene aggiornato il contatore 
			isParola = false	//e ri-iniziallizata la variabile booleana isParola
		}						
		// altrimenti se isParola è false non succede niente e si va avanti (questo serve nel caso la frase non iniziasse con una lettera)
	}

	if isParola{		//se dopo aver terminato il ciclo for la variabile isParola è ancora true significa che la frase è finita con una parola
		n++		//che non è stata ancora contata (perchè la frase in sè non finisce con un carattere speciale ) e quindi la aggiungo al contatore
	}								                  	

	return			//return del contatore
}
