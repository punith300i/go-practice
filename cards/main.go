package main

import "fmt"

func main(){
	cards := []string{ "Ace of Diamonds", newCard()}
	cards = append(cards, "Another Card")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	
}

func newCard() string {
	return "5 of Diamonds"
}