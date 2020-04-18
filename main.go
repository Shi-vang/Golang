package main

import "fmt"

func main() {
	//var card string = "Ace of spade"
	splic1 := newDeck()
	splic2 := newDeck()
	splic1.print()
	fmt.Println("  ")
	splic2.print()
}
