package main

import (
	"fmt"
)

/*	Dati due vettori di interi, stampare i numeri che sono presenti in entrambi gli array	*/

func main(){
	slice1 := []int{1, 2, 3, 4, 8, 5, 4}		//slice1
	slice2 := []int{3, 4, 5, 5, 5, 6}			  //slice2

	fmt.Println(comune(slice1, slice2))			//stampa elementi comuni
}

func comune(num1, num2 []int) (com []int){		//funzione che prende due slice di interi e restituisce una slice che contiene gli elementi comuni
	for _, n1 := range num1 {					          //for range sulla slice1
		for _, n2 := range num2 {			          	//for range sulla slice 2
			if n1 == n2 {				                		//se slice1[x] == slice2[y]
				flag := false				                	//variabile booleana flag inizializzata a false se il numero della slice1 è uguale al numero della slice2
				for _, v := range com {			          //for sulla slice che contiene gli elementi comuni
					if n1 == v {				                //se l'elemento della slice1 è già presente nella slice di elementi comuni 
						flag = true				                //flag = true (significa che il numero corrente è già presente nell'array degli elementi comuni)
						break				                    	//esco dal ciclo perchè flag = true
					}
				}

				if !flag{						                  //se flag = false (ovvero se il numero corrente non è presente nell'array degli elementi comuni)
					com = append(com, n1)		            //aggiungo il numero nella slice di elementi comuni
					break						                    //esco dal ciclo (perchè so che il numero corrente è presente nella slice2)
				}
			}
		}
	}
	return
}
