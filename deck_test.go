package main

import (
	"testing"
)

func TestNewDeckFromFile(t *testing.T) {
	deck := newDeckFromFile("cards.txt")

	if len(deck) != 4 {
		t.Errorf("Error, Expected length 5 got %v", len(deck))
	}
}
