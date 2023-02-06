package main

import (
	"fmt"
	"os"
)

func main(){
	m := subString(os.Args[1])
	if len(m) != 0{
		for k, v := range m {
			fmt.Println(k, " -> occ: ", v)
		}
	}
}

func subString(s string) (m map[string]int){
	m = make(map[string]int)
	for start := 0; start < len(s); start++ {
		for end := start+3; end <= len(s); end++ {
			num := s[start:end]
			if num[0] == num[len(num)-1]{
				m[num]++
			}
			
		}
	}
	return 
}