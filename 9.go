package main

import (
	. "fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main(){
	num, err := strconv.Atoi(os.Args[1])
	if err != nil{
		Println("inserted value not valid")
		os.Exit(1)
	}
	generation := generate(num)

	Println("generated values:", generation)
	Println("valued above soglia:", generation[0:len(generation)-1])
}

func generate(soglia int) (numbers[]int){
	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < soglia; i++ {
		generatedNumber := rand.Intn(101)
		if generatedNumber < soglia {
			break
		}
		numbers = append(numbers, generatedNumber)
	}
	return
}