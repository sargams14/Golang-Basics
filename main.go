package main

func main() {
	cards := newDeck()
	/* 	hand, remainingDeck := deal(cards, 5)
	   	hand.print()
	   	remainingDeck.print() */
	//fmt.Println(cards.saveToFile("deck_of_cards.txt"))
	//cardsDeck := newDeckFromFile("deck_of_cards.txt")
	cards.shuffle()
	cards.print()
}
