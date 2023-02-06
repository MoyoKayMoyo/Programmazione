package main

import (
    . "fmt"
    "strconv"
)

func main()  {
    var n, b uint
    Print("insert a number (base 10): ")
    Scan(&n)
    Print("insert the base u wanna covert in: ")
    Scan(&b)
    
    s, e := base10toX(n, b)
	if e {
		Println(n, "in base", b, "is", s)
	} else {
		Println("the base", b, "is not valid")
	}
}

func base10toX(n, b uint) (string, bool){
    if b < 2 || b > 16 {
        return "", false
    }

    var s string
    lsb := n%b
    for n != 0 {
        lsb = n%b
        if lsb >= 10 && lsb <= b {
            switch lsb {
            case 10:
                s = "A" + s
            case 11:
                s = "B" + s
            case 12:
                s = "C" + s
            case 13:
                s = "D" + s
            case 14:
                s = "E" + s
            case 15:
                s = "F" + s    
            }

        } else if lsb < 10 {
            s = strconv.Itoa(int(lsb)) + s
        } else if lsb > b {
            num := b-10 
            s = strconv.Itoa(int(num)) + s
        } else if b < 10 {
            for lsb != 0 {
                s = strconv.Itoa(int(lsb)) + s
            }
        }
        n /= b
    }

    return s, true
}