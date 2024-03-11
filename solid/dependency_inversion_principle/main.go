// This example demonstrates Dependecy Inversion Principle (DIP), one of the
// SOLID design principles (`D` letter). SOLID is a mnemonic acronym for five
// design principles intended to make object-oriented designs more
// understandable, flexible, and maintainable.
//
// DIP says:
//  1. High Level Modules (business logic) should not depend on Low LM
//     (storage, network, etc.).
//  2. Both should depend on abstractions (abstract classes or interfaces).
package main

// Main function to demonstrate DIP in action
func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate()
}
