package main

func main() {
	// var card string = "Ace of Spades"
	cards := newDeck()
	cards.saveToFile("my_cards")

}
