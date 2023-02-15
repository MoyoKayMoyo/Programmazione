package main

import "fmt"

func main() {

	s := "filtéo"
	sc := []rune(s)
	for i:=0;i<len(s);i++{
		inversa := string(sc[i:])+string(sc[:i])
		fmt.Println(inversa)
	}
}
