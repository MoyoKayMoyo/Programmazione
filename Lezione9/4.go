package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){

	if len(os.Args) > 1{
		fmt.Println(Sottostringhe(os.Args[1]))
		
	}
}

func Sottostringhe(s string) (result []string){
	_, err := strconv.Atoi(s)
	if err != nil{
		return
	}

	var num string
	for start := 0; start < len(s); start++ {
		for i := start+2; i <= len(s); i++ {
			num = s[start:i]
			if valid(num){
				result = append(result, num)
			}
		}
	}
	return
}

func valid(num string) bool{
	for k := 0; k < len(num)-1; k++ {
		if num[k] > num[k+1]{
			return false
		}
	}
	return true
}