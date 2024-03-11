// This example demonstrates Single Responsibility Principle (SRP), one of the
// SOLID design principles (`S` letter). SOLID is a mnemonic acronym for five
// design principles intended to make object-oriented designs more
// understandable, flexible, and maintainable.
//
// SRP says: "There should never be more than one reason for a class to change.
// In other words, every class should have only one responsibility".
package main

import (
	"fmt"
)

// Main function to demonstrate SRP in action
func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	fmt.Println(j.String())

	//
	SaveToFile(&j, "journal.txt") // right approach

	//
	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt") // right approach
}
