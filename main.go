package main

func main() {
	cards := deck{"Card", newCard(), "Another card"}
	cards = append(cards, "Fourth card")

	cards.SaveToFile("cards.txt")
}

func newCard() string {
	return "New card"
}
