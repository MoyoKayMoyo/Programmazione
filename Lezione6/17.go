package main

import . "fmt"

func main(){
	var c rune 	= 127153 // '\U0001F0B1'
	n := 10

	for i := 0; i < n; i++ {
		Println("Symbol: ", string(c), " - Numeric code in base 10:", c)
		c++
	}

}