package main

func main() {
	// declare new slice
	cards := deck{"Two of Ace", newCard()}

	// append to a slice
	cards = append(cards, "Six of spades")

	cards.print()
}

func newCard() string {
	return "The new card is five of diamonds"
}
