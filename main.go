package main

func main() {
	cards := newDeck()
	cards = append(cards, newDeck()...)
	cards.print()
}
