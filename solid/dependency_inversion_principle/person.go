// Person
package main

// Struct for a person with a name
type Person struct {
	name string
}

// Struct to hold information about a relationship between two people
type Info struct {
	from         *Person
	to           *Person
	relationship Relationship
}
