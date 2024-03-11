// Research
package main

import "fmt"

// High-level module that uses a RelationshipBrowser to perform research
type Research struct {
	//relationships Relationships // this approach breaks DIP
	browser RelationshipBrowser // this approach doesn NOT break DIP
}

// This method is commented out, because it showcases a direct dependency on the
// Relationships struct, against the Dependency Inversion Principle (DIP). It
// directly accesses the `relations` field, seeking to find and announce "John"'s
// children by sifting for the Parent relationship. This tight coupling with
// Relationships' implementation details undermines code flexibility and
// maintainability. It breaches DIP by depending on a concrete class rather than
// an abstraction. An alternate DIP-compliant method using the
// RelationshipBrowser interface replaces this, enabling the high-level module
// (Research) to stay decoupled from the low-level module (Relationships),
// aligning with DIP's advocacy for module dependency on abstractions.
// func (r *Research) Investigate() {
// 	relations := r.relationships.relations
// 	for _, rel := range relations {
// 		if rel.from.name == "John" &&
// 			rel.relationship == Parent {
// 			fmt.Println("John has a child called ", rel.to.name)
// 		}
// 	}
// }

// Uses the RelationshipBrowser to find and print all children of "John"
func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called ", p.name)
	}
}
