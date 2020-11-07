package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	cards := newDeck()
	// cards = append(cards, "This is a new element")

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
	greeting := "hi there"
	fmt.Println([]byte(greeting))
}
