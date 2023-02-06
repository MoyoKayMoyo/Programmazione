package main

import . "fmt"

func main() {
	Printf("Inserisci l'operazione da effettuare\n")

	Println("1: Secondi -> ore")
	Println("2: Secondi -> minuti")
	Println("3: Minuti -> ore")
	Println("4: Minuti -> secondi")
	Println("5: Ore -> secondi")
	Println("6: Ore -> minuti")
	Println("7: Minuti -> giorni, ore")
	Println("8: Minuti -> anni, giorni")

	// anno bisestile int
	// data int int int d m y valida
	var scelta, val int
	Print(" : ")
	Scan(&scelta)
	if scelta < 1 || scelta > 8 {
		Println("Scelta non valida")
	} else {
		Print("Inserire il valore da convertire: ")
		Scan(&val)

		if scelta == 1 {
			Println(val, "secondi equivalgono a", float64(val)/3600, "ore")
		} else if scelta == 2 {
			Println(val, "secondi equivalgono a", float64(val)/60, "minuti")
		} else if scelta == 3 {
			Println(val, "minuti equivalgono a", float64(val)/60, "ore")
		} else if scelta == 4 {
			Println(val, "minuti equivalgono a", float64(val)*60, "secondi")
		} else if scelta == 5 {
			Println(val, "ore equivalgono a", float64(val)*3600, "secondi")
		} else if scelta == 6 {
			Println(val, "ore equivalgono a", float64(val)*60, "minuti")
		} else if scelta == 7 {
			var p int
			p = (val / 60) / 24
			Println(val, "minuti equivalgono a", (val/60)/24, "giorni e ", ((float64(val)/60)/24-float64(p))*24, "ore")
		} else if scelta == 8 {
			Println(val, "minuti equivalgono a", ((val/60)/24)/365, "anni e", (val/60)/24, "giorni")
		}
	}

}
