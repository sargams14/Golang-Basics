package main

func main() {

	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, newDeck()...)
	cards.print()
}

func newCard() string {
	return "Two of Hearts"
}
