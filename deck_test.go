package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 3 {
		t.Errorf("Error, Expected length 3 got %v", len(deck))
	}

	if deck[0] != "One" {
		t.Errorf("Error, Expected One, got %v", deck[2])
	}

	if deck[(len(deck)-1)] != "Three" {
		t.Errorf("Error, Expected Three, got %v", deck[2])
	}
}

func TestNewDeckFromFile(t *testing.T) {
	os.Remove(GetFileName())

	deck := newDeck()
	deck.saveToFile(GetFileName())
	deckFromFile := newDeckFromFile(GetFileName())

	if len(deckFromFile) != 3 {
		t.Errorf("Error, Expected length 3 got %v", len(deck))
	}

	os.Remove(GetFileName())
}

func newDeck() deck {
	return deck{"One", "Two", "Three"}
}

func GetFileName() string {
	return "_testingdeck"
}
