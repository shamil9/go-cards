package main

import (
	"fmt"
)

func main() {
	cards := deck{"Card", newCard(), "Another card"}
	cards = append(cards, "Fourth card")

	hand, remainingCards := deal(cards, 2)

	fmt.Println(hand.ToString(), remainingCards.ToString())
}

func newCard() string {
	return "New card"
}
