package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
	
	// call the receiver_functions function and print the result
	fmt.Println(receiver_functions())
}

func newCard() string {
	return "The new card is five of diamonds"
}
