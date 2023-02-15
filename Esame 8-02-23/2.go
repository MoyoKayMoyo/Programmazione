package main

import "fmt"

func main(){
	
	fmt.Println(Romano2decimale("MCMLXXXIII"))
}

func Romano2decimale(ro string) (dec int){
	numeriRomani := map[string]int{"I":1, "V":5, "X":10, "L":50, "C":100, "D":500, "M":1000}
	for i := 0; i < len(ro)-1; i++ {
		if numeriRomani[string(ro[i])] >= numeriRomani[string(ro[i+1])]{
			dec += numeriRomani[string(ro[i])]
		} else {
			dec += numeriRomani[string(ro[i+1])] - numeriRomani[string(ro[i])]
			i++
		}
	}
	dec += numeriRomani[string(ro[len(ro)-1])]

	return
}