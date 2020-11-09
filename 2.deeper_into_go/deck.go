package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

// This means we can call the deal function with:
// First agument of type deck, second argument of type int
// We cannot call the deal function with any other type of function.
// If we had fmt.Println(d) it will print out the entire deck "object"
//  This returns two values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// This takes a slice of strings and converts it to one string
	return strings.Join([]string(d), ",")
}

// Convert the string to byte slice and convert it to a file

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
}
