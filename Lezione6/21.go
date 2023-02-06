package main

import . "fmt"

func main() {
	var order = -1
	var total, amount int
	Println("What do u want to order?")

	for order != 0 {
		Println("\t1. French fries 2$")
		Println("\t2. Hamburger 5$")
		Println("\t3. Coca cola 2$")
		Println("\t0. Termina")
		Print("Order: ")
		Scan(&order)

		switch order {
		case 1:
			Print("French fries? how many?: ")
			Scan(&amount)
			total += 2 * amount
		case 2:
			Print("Hamburger? how many?: ")
			Scan(&amount)
			total += 5 * amount
		case 3:
			Print("Coca cola? how many?: ")
			Scan(&amount)
			total += 2 * amount
		}

	}

	Println("Thats", total, "$ + 2 for shipping. Total", total+2)

}
