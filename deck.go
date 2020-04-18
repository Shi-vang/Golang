package main

import "fmt"

// it is kind of we typed string type to deck
type deck []string

// this function bound with deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	card := deck{"jack"}
	suitCard := []string{"Spade", "Diamond", "club", "heart"}
	valuesCard := []string{"Ace", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "king", "Queen"}

	for _, suit := range suitCard {
		for _, value := range valuesCard {
			card = append(card, value+" of "+suit)
		}
	}
	return card
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
