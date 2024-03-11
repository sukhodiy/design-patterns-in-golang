// This example demonstrates Open-closed Principle (OCP), one of the
// SOLID design principles (`O` letter). SOLID is a mnemonic acronym for five
// design principles intended to make object-oriented designs more
// understandable, flexible, and maintainable.
//
// OCP says: "Software entities should be open for extension, but closed for
// modification".
//
// This example contains:
// 1. Bad filter
// 2. Specification pattern
// 3. Idiomatic filter
package main

import "fmt"

// Main function to demonstrate OCP in action
func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products (old):\n")
	f := BadFilter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Printf("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Printf("Large green products:\n")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}

	idiomf := IdiomaticFilter{}
	prs := idiomf.Filter(products, func(p *Product) bool {
		return p.color == green && p.size == large
	})
	fmt.Printf("Large green products (idiomatic):\n")
	for _, v := range prs {
		fmt.Printf(" - %s is large and green\n", v.name)
	}

	prs2 := IdiomaticFuncFilter(products, func(p *Product) bool {
		return p.color == green && p.size == large
	})
	fmt.Printf("Large green products (idiomatic func):\n")
	for _, v := range prs2 {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
