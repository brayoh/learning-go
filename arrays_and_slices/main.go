package main

import "fmt"

func main() {
	// declare new slice
	// A slice is a dynamically-sized, flexible view into the elements of an array.
	// A slice is a descriptor of an array segment. 
	// It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).
	// A slice does not store any data, it just describes a section of an underlying array.
	cards := []string{"Two of Ace", newCard()}

	// append to a slice
	// An append operation creates a new slice with the new value added to the end of the original slice. 
	// The original slice is not modified.
	cards = append(cards, "Six of spades")

	fmt.Println(cards)

	// Iterate through a slice
	// The range form of the for loop iterates over a slice or map.
	// When ranging over a slice, two values are returned for each iteration. 
	// The first is the index, and the second is a copy of the element at that index.
	// If you don't need the index, you can ignore it with the blank identifier _.
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "The new card is five of diamonds"
}
