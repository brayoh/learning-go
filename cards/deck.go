package main

import "fmt"

// Create a new type of deck
// which a slice of strings
type deck []string


func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Flowers", "Aces", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for _, suit := range cardSuits {
		for _, value := range cardValues { // the underscore is used to ignore the index and only use the value
			// append to the slice
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

// Receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
