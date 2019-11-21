package main

import "fmt"

// Create a new type of deck
// which a slice of strings
type deck []string

// Receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
