package main

import . "fmt"

func main() {
	var h, m, s int
	Print("Inserire i secondi:")
	Scan(&s)

	m = s/60 

	if m%60 == 0 {
		h = m/60
	}

	Println(h, ":", m, ":", s)
}
