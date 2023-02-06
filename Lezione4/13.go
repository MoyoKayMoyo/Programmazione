package main

import . "fmt"

func main(){
	var num, num_f int
	Print("Inserire un numero: ")
	Scan(&num)

	prec_1 := 1	
	prec_2 := 0
	num_f = 1
	Println(num_f)

	for i := 1; i < num; i++ {
		prec_2 = prec_1
		prec_1 = num_f
		if i == 1 {
			num_f = 1
		} else {
			num_f = prec_1 + prec_2
		}
		Println(num_f)
	}

}