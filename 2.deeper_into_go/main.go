package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of Spades"
	cards := []string{"Ace of diamonds", newCard()}
	cards = append(cards, "This is a new element")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string { // Indicates this function will return a string
	return "Five of Diamonds"
}
