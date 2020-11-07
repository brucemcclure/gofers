package main

import "fmt"

// Create a new type of deck
//  Which is a slice of strings
// deck === []string
type deck []string

// Anytime newDeck is called it returns an 'instance' of deck
func newDeck() deck {
	cards := deck{} // Creates a new slice varibale called cards

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards // returning the cards variable that is actully of type deck
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
