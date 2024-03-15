package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
	fmt.Println(cards.toString())
	cards.saveToFile("test.txt")
	var newCards deck = newDeckFromFile("test.txt")
	newCards.print()
	newCards.shuffle()
	fmt.Println("Test")
	newCards.print()
}

func newCard() string {
	return "5 of Diamonds"

}
