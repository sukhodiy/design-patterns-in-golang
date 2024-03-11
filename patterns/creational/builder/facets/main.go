package main

import "fmt"

// Example: composite builder for person's info.

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostcode("SW12BC").
		Works().
		At("Fabrikam").
		AsA("Programmer").
		Earning(123000)
	person := pb.Build()
	fmt.Println(person)
}
