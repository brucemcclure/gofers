package main

func main() {
	// var card string = "Ace of Spades"
	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "This is a new element")

	cards.print()

}

func newCard() string { // Indicates this function will return a string
	return "Five of Diamonds"
}
