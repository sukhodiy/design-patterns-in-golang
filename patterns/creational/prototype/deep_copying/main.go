package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"},
	}

	// jane := john       // copy pointer for Address
	// jane.Name = "Jane" // ok
	// jane.Address.StreetAddress = "321 Baker St"

	// deep copying - approach to copy values, not pointers

	jane := john
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country,
	}
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
