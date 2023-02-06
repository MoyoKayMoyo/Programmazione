package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main(){
	suff, ins := filterMarka(readNumbers())
	Println("good marks: ", suff)
	Println("bad marks: ", ins)

}
func readNumbers() []int{
	var numbers[]int
	for _, v := range os.Args [1:]{
		num, err := strconv.Atoi(v)
		if err == nil && num >= 0 &&  num <= 100{
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func filterMarka(voti[]int) (suff, ins []int){
	for _, v := range voti {
		if v < 60 {
			ins = append(ins, v)
		} else {
			suff = append(suff, v)
		}
	}
	return
}