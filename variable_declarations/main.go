package main

import "fmt"

func main() {
	// explicit declarations
	var card string = "Ace of spades"

	// implicit declarations - go compiler infers the type
	card2 := "Another ace of spades"

	// re-assign value
	card2 = "Another ace of spades that's been re-assigned"

	fmt.Println("First card is", card)
	fmt.Println("Second card is", card2)
}
