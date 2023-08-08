package main

import "fmt"

// Custom type
type deck []string

// Using a receiver
// Any variable of type deck can access this method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {

	return deck{"Four of Diamonds", "Five of Hearts", "Eight of Spades", "Ace of Clubs", "King of Hearts"}
}
