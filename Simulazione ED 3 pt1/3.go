package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)
type Punto struct{
	etichetta string
	ascissa float64
	ordinata float64
}

func main(){
	tragitto := nuovoTragitto()
	if len(tragitto) < 2{
		fmt.Println("inserisci almeno 2 punti")
	} else {
		fmt.Printf("Lunghezza percorso: %.2f\n", Lunghezza(tragitto))
	}
}

func nuovoTragitto()(tragitto []Punto){
	var punto Punto
	var coordinate []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		coordinate = strings.Split(txt, ";")
		punto.etichetta = coordinate[0]

		num, err := strconv.ParseFloat(coordinate[1], 1)
		if err == nil{
			punto.ascissa = float64(num)
			num, err = strconv.ParseFloat(coordinate[2], 1)
			if err == nil{
				punto.ordinata = float64(num)
			} else {
				fmt.Println("ordinata insertita non è valida")
				break
			}
			fmt.Println(String(punto))
			tragitto = append(tragitto, punto)
		} else {
			fmt.Println("ascissa insertita non è valida")
			break
		}
	}
	return
}

func tratta(p1, p2 Punto) (d float64){
	// ((x1-x2)^2 + (y1-y2)^2)^1/2
	return math.Sqrt(math.Pow(p1.ascissa-p2.ascissa, float64(2)) + math.Pow(p1.ordinata-p2.ordinata, float64(2)))
}

func String(p Punto) string {
	return p.etichetta + ": (" + fmt.Sprintf("%.1f",p.ascissa) + " " + fmt.Sprintf("%.1f",p.ordinata) + ")"
}

func Lunghezza(tragitto []Punto) (d float64){
	for k := 0; k < len(tragitto)-1; k++ {
		d += tratta(tragitto[k], tragitto[k+1])
		fmt.Println(d)
		fmt.Printf("distanza tra %s e %s = %.1f\n", tragitto[k].etichetta,tragitto[k+1].etichetta, tratta(tragitto[k], tragitto[k+1]))
	}
	return 
}