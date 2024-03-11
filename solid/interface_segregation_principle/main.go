// This example demonstrates Interface Segregation Principle (ISP), one of the
// SOLID design principles (`I` letter). SOLID is a mnemonic acronym for five
// design principles intended to make object-oriented designs more
// understandable, flexible, and maintainable.
//
// ISP says: "Clients should not be forced to depend upon interfaces that they
// do not use".
package main

// Main function to demonstrate ISP in action
func main() {
	doc := Document{}

	var wm WrongMachine // bad interface type
	wm = &MultiFunctionPrinter{}
	wm.Fax(doc)

	var pr Printer // good interface type
	pr = &MyPrinter{}
	pr.Print(doc)
}
