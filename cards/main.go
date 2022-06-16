package main

func main() {

	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// fmt.Println()
	// remainingDeck.print()

	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
