package main

import (
	"fmt"
	"os"
	"strings"
)

// Custom type
type deck []string

// Using a receiver, Any variable of type deck can access this method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// ioutil is deprecated since go 1.16, so we use os instead
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	fileByte, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error::", err)
		os.Exit(1)
	}
	return deck(strings.Split((string(fileByte)), ","))
}
