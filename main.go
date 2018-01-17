package main

func main() {
	cards := deck{"Card", newCard(), "Another card"}
	cards = append(cards, "Fourth card")

	hand, remainingCards := deal(cards, 2)

	hand.print()
	remainingCards.print()
}

func newCard() string {
	return "New card"
}
