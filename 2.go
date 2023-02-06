package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main(){
	valid := true
	var seq []int

	for i, v := range os.Args[1:]{
		num, err := strconv.Atoi(v)

		if err == nil {
			seq = append(seq, num)
		} else {
			Println("element in position", i, "not valid")
			valid = false
			break
		}
	}

	if valid {
		for i := 0; i < len(seq); i++ {
			if i%2 == 1 && seq[i] >= seq[i-1] || i%2 == 0 && i != 0 && seq[i] <= seq[i-1]{
				Println("element in position", i, "not valid")
				valid = false
				break
			}
		}
	}

	if valid {
		Println("the sequence is valid")
	}
}