package main

import "fmt"

// type deck : slice of strings
type deck []string

func newDeck() deck{
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue + " of " + cardSuit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print(){
	for i, card := range d {
		fmt.Println(i,card)
	}
}

