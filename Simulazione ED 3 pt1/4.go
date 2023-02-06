package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	num := os.Args[1]
	diff := os.Args[2]
	var n string
	if len(num) > len(diff) {
		d, _ = strconv.Atoi(os.Args[2])
		for j := 0; j < len(num)-1; j++ {

			for i := d; i < len(num)-1; i++ {
				n = 
				for k := i+1; k < len(num); k++ {
					fmt.Println(string(num[i]), string(num[k]))
				}
			}
		}
		
			
	} else {
		fmt.Println("Il primo numero deve essere maggiore del secondo")
	}

}