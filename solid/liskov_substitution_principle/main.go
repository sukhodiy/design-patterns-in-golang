// This example demonstrates Liskov Substitution Principle (LSP), one of the
// SOLID design principles (`L` letter). SOLID is a mnemonic acronym for five
// design principles intended to make object-oriented designs more
// understandable, flexible, and maintainable.
//
// LSP says: "Functions that use pointers or references to base classes must
// be able to use objects of derived classes without knowing it".
package main

import "fmt"

// 1. Set height for figure equal 10
// 2. Calculate and print expected and actual area for figure
func UseIt(sized Sized) {
	height := 10
	width := sized.GetWidth()

	sized.SetHeight(height)

	expectedArea := height * width
	actualArea := sized.GetWidth() * sized.GetHeight()

	fmt.Println("Expected area =", expectedArea,
		", actual area =", actualArea)
}

// Main function to demonstrate LSP in action
func main() {
	rc := &Rectangle{2, 3}
	fmt.Println("Rectangle{2, 3}")
	// Correct results
	UseIt(rc) // Output: Expected area = 20 , actual area = 20

	fmt.Println("NewSquare(5)")
	sq := NewSquare(5)
	// Incorrect results
	UseIt(sq) // Output: Expected area = 50 , actual area = 100
}
