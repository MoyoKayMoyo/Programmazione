package main

import (
	"errors"
	. "fmt"
	"math/rand"
)

type GameCard struct{
	seed string 
	symbol string
}

type Deck struct{
	cards []GameCard
	numCards int
}

// gestire il drawn di più carte

func main(){
	var drawn_card []GameCard
	deck := createDeck()

	mixDeck(deck.cards)

	last_card, deck, err := drawCard(deck)
	drawn_card = append(drawn_card, last_card)
	if err == nil{
		Println("drawn card: ", drawn_card[len(drawn_card)-1])
		printDeck("\ndraw deck", deck)
	} else {
		Println(err)
	}
	
	last_card, deck, err = drawCard(deck)
	drawn_card = append(drawn_card, last_card)
	if err == nil{
		Println("drawn card: ", drawn_card[len(drawn_card)-1])
		printDeck("\ndraw deck", deck)
	} else {
		Println(err)
	}

	last_card, deck, err = drawCard(deck)
	drawn_card = append(drawn_card, last_card)
	if err == nil{
		Println("drawn card: ", drawn_card[len(drawn_card)-1])
		printDeck("\ndraw deck", deck)
	} else {
		Println(err)
	}

	if len(drawn_card)-1 >= 0{
		deck, err = returnCard(deck, drawn_card[len(drawn_card)-1])
		if err == nil {
			Println("card: ", drawn_card[len(drawn_card)-1])
			drawn_card = drawn_card[:len(drawn_card)-1]
			printDeck("\nreturn card deck:", deck)
		} else {
			Println(err)
		}
	} else {
		Println("there are no card to put back in the deck")
	}	
		
}

func createCard(seed, symbol string) GameCard{
	return	GameCard{seed: seed, symbol: symbol}
}

func createDeck() Deck{
	seed := []string{"♦", "♠", "♣", "♥"}
	symbol := []string{"A", "2", "3", "4", "5", "6", "7", "J", "Q", "K"}
	var deck Deck
	for _, k := range seed {
		for _, v := range symbol {
			deck.cards = append(deck.cards ,createCard(k,v))
			deck.numCards++
		}
	}
	return deck 
}

func mixDeck(cards []GameCard){
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
}

func drawCard(deck Deck)(GameCard, Deck, error){
	if deck.numCards > 0 {
		first := deck.cards[len(deck.cards)-1]
		deck.cards = deck.cards[:len(deck.cards)-1]
		deck.numCards--
	
		return first, deck, nil
	}
	
	return deck.cards[0], deck, errors.New("deck is empty")
}

func returnCard(deck Deck, last GameCard) (Deck, error){
	if deck.numCards <= 40{
		deck.cards = append(deck.cards, last)
		deck.numCards++
		return deck, nil
	}
	return deck, errors.New("deck is full")
}

func printDeck(text string, deck Deck){
	Println(text)
	for i := 0; i < len(deck.cards); i=i+2{
		if i%20 != 0 && i+1 != len(deck.cards){
			Print(deck.cards[i], " ", deck.cards[i+1])
		} else if i+1 == len(deck.cards){
			Print(deck.cards[i])
		}
		Println()
	}
}
