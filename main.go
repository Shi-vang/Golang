package main

import "fmt"

func main() {
	//var card string = "Ace of spade"
	splic1 := newDeck()
	splic2 := newDeck()
	hand1, remainhand1 := deal(splic1, 3)
	splic1.print()
	fmt.Println(hand1, remainhand1)
	splic1.saveToFile("splice1 file")
	fmt.Println("  ")
	hand2, remainhand2 := deal(splic2, 7)
	fmt.Println(hand2, remainhand2)
	splic2.print()
	splic2.saveToFile("splice2 file")
}
