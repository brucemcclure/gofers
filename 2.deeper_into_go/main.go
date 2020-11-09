package main

func main() {
	// var card string = "Ace of Spades"
	// Creating a new deck
	cards := newDeck()
	cards.saveToFile("my_cards")

	// Loading up a new deck
	// cards := newDeckFromFile("my_cards")
	cards.print()
}
