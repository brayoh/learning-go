package main

import "fmt"

// Create a new type of deck
// which a slice of strings
type deck []string


func newDeck() deck {
	cards := deck{}

	// create a slice of strings
	cardSuits := []string{"Spades", "Flowers", "Aces", "Diamonds", "Hearts", "Clubs"}

	// create a slice of strings
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	// create a nested loop to iterate over the cardSuits and cardValues slices
	for _, suit := range cardSuits {
		for _, value := range cardValues { // the underscore is used to ignore the index and only use the value
			// append to the slice
			cards = append(cards, suit+" of "+value)
		}
	}
	
	return cards
}

// Receiver function
// This function is defined on the deck type and can be called on any instance of deck
// The receiver function is defined with the receiver type (d deck) and the function name (print)
// The receiver function can access the fields and methods of the deck type and can be called on any instance of deck
// The receiver function can be called on any instance of deck and can be used to print the contents of the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// return two slices
	// the first slice is the hand of cards, and the second slice is the remaining cards in the deck
	return d[:handSize], d[handSize:]
}
