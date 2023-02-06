package main

import (
    . "fmt"
    "os"
    "strconv"
)

func main(){
    num, err := strconv.Atoi(os.Args[1])
    if err == nil {
        Println("fattoriale: ", fatt(num))
    } else {
        Println("not valid")
    }
}

func fatt(n int)(f []int){
    v := 1
    for i := 1; i <= n; i++ {
        if i == 1 {
            f = append(f, 1)
        } else {
            v *= i-1
            f = append(f, i*v)
        }
    }
    return 
}