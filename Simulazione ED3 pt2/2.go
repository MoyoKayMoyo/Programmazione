package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	if len(os.Args) < 2{
		fmt.Println("-1")
	} else {
		var numeri []uint
		var err error
		var num int
		for _, n := range os.Args[1:] {
			num, err = strconv.Atoi(n)
			if err == nil {
				numeri = append(numeri, uint(num))
			}
		}
		fmt.Println(mcm_sequenza(genera(contaOccorenze(numeri))))
	}
}

func contaOccorenze(numeri []uint) map[uint]uint{
	m := make(map[uint]uint)
	for _, v := range numeri {
		m[v]++
	}

	return m
}

func genera(occorrenze map [uint]uint) (num []uint){
	for k, v := range occorrenze {
		if v == 1 {
			num = append(num, k)
		}
	}

	return
}

func MCD(a, b uint) uint{
	for b!= 0 {
		a, b = b, a%b
	}
	return  a
}

func mcm(a, b uint) uint{
	return (a*b)/MCD(a, b)
}

func mcm_sequenza(num []uint) int {
	if len(num) < 2 {
		return -1
	}
	
	p_mcm := mcm(num[0], num[1])
	for i := 2; i < len(num); i++ {
		p_mcm = mcm(p_mcm, num[i])
	}

	return 	int(p_mcm)
}