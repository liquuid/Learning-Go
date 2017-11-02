package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs, but got %v", d[len(d)-1])
	}

}
func TestSaveDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktestfile.txt")

	deck := newDeck()

	deck.saveToFile("_decktestfile.txt")
	loadedDeck := newDeckFromFile("_decktestfile.txt")

	if len(loadedDeck) != 16 {
		t.Errorf("Expect 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktestfile.txt")

}
