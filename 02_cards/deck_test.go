package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of 16, but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Error("Expected first card of Ace of spades")
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadDeck := newDeckFromFile("_decktesting")
	if len(loadDeck) != 16 {
		t.Error("Expected first card of Ace of spades")
	}
	os.Remove("_decktesting")
}
