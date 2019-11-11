package main

import "fmt"

func main() {
	// declare new slice
	cards := []string{"Two of Ace", newCard()}

	// append to a slice
	cards = append(cards, "Six of spades")

	fmt.Println(cards)

	// loop through a slice

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "The new card is five of diamonds"
}
