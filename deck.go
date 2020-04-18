package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(data), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < len(d); i++ {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
