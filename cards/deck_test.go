package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected length is wrong %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("First card needs to Ace")
	}
}

func TestSaveToDeck(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")
	newDeck := newDeckFromFile("_decktesting")

	if len(newDeck) != 12 {
		t.Errorf("The length does not match %v", len(newDeck))
	}

	os.Remove("_decktesting")

}
