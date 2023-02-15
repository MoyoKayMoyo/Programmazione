package main

import "fmt"

func main(){

	s := "cartaèluuunga"
	max := 0
	maxString := ""
	for start := 0; start < len(s); start++ {
		for i := start; i <= len(s); i++ {
			num := s[start:i]
			if onlyAscii(num){
				if max < len(num){
					max = len(num)
					maxString = num
				}
            }
		}
	}

	fmt.Println(maxString)
}

func onlyAscii(s string) bool{
	for _, c := range s {
		if len(string(c)) > 1 {
			return false
		}
	}
	return true
}