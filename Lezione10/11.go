package main

import (
	. "fmt"
	"os"
)

func main(){
    s := os.Args[1]
    var m map[string]int
    m = make(map[string]int)
    for start := 0; start < len(s); start++ {
		for i := start+3; i <= len(s); i++ {
			num := s[start:i]
			if valid(num){
                m[num]++
            }
		}
	}

    if len(m) != 0{
        for k, v := range m {
            Println(k, " -> occ: ", v)
        }
    }

}

func valid(s string) bool{
    if s[0] == s[len(s)-1]{
        return true
    }

    return false
}

