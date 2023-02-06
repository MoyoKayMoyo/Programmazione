package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Evento struct{
	inizio string
	fine string
	nome string
}

func main(){
	agenda := nuovaAgenda()
	if len(agenda) < 2 {
		fmt.Println("moyo, need more than 1")
	} else {
		for i := 0; i < len(agenda); i++ {
			for k := i; k < len(agenda); k++ {
				if !isPrecedente(agenda[i], agenda[k]){
					agenda[i], agenda[k] = agenda[k], agenda[i]
				}
			}
		}

		for _, v := range agenda {
			fmt.Println(StringEvento(v))
		}

		fmt.Println()

		fmt.Println(StringEvento(agenda[0]))
		for i := 0; i < len(agenda)-1; i++ {
			if !isSovrapposto(agenda[i], agenda[i+1]) {
				fmt.Println(StringEvento(agenda[i+1]))
			}
		}
	}
}

func nuovaAgenda() (agenda[]Evento){
	var e Evento
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		text := scanner.Text()
		moyo := strings.Split(text, " ")
		inizio := strings.Split(moyo[0], ":")
		fine := strings.Split(moyo[1], ":")
		nome := moyo[2]

		if isValid(inizio) && isValid(fine){
			e.inizio = inizio[0]+":"+inizio[1]
			e.fine = fine[0]+":"+fine[1]
			e.nome = nome
			agenda = append(agenda, e)
		}
	}
	return 
}

func isValid(time []string) bool{
	h, err_h := strconv.Atoi(time[0])
	if err_h == nil && h >= 0 && h <= 23{
		m, err_m := strconv.Atoi(time[1])
		if err_m == nil && m >= 0 && m <=59{
			return true
		}
		return false
	}
	return false
}

func StringEvento(e Evento) string {
	return e.inizio + "-" + e.fine + " " + e.nome
}

func isPrecedente(e1, e2 Evento) bool {
	inizio1 := strings.Split(e1.inizio, ":")
	inizio2 := strings.Split(e2.inizio, ":")

	if inizio1[0] < inizio2[0] {
		return true
	} else if inizio1[0] > inizio2[0] {
		return false
	} else {
		fine1 := strings.Split(e1.fine, ":")
		fine2 := strings.Split(e2.fine, ":")

		if inizio1[1] <= inizio2[1]{
			return true
		} else if inizio1[1] == inizio2[1] && fine1[0] <= fine2[0]{
			return false
		}
		return false
	}
}

func isSovrapposto(e1, e2 Evento) bool {
	fine1 := strings.Split(e1.fine, ":")
	inizio2 := strings.Split(e2.inizio, ":")

	if inizio2[0] == fine1[0]{
		return true
	}
	return false
}