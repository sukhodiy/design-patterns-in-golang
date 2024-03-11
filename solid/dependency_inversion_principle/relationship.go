// Relationship
package main

// Enumeration for different types of relationships
type Relationship int

// Constants representing the type of relationships
const (
	Parent Relationship = iota
	Child
	Silbling
)

// Interface that defines the method to find all children of a given person
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// Stores relationships and implements RelationshipBrowser
type Relationships struct {
	relations []Info
}

// Method to find all children of a given person by their name
func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationship == Parent &&
			v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}

	return result
}

// Method to add a parent-child relationship
func (r *Relationships) AddParentAndChild(
	parent, child *Person) {
	r.relations = append(r.relations,
		Info{parent, child, Parent})
	r.relations = append(r.relations,
		Info{child, parent, Child})
}
