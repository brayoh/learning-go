package main

func main() {
	// declare new slice
	cards := newDeck()

	// append to a slice
	cards = append(cards, "Six of spades")

	hand, _ := deal(cards, 5)

	// print the cards
	hand.print()
}