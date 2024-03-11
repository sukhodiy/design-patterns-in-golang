package main

import (
	"fmt"
)

// Example: can a person drive a car because of his age?

func main() {
	p := NewPerson(15)
	t := &TrafficManagement{p.Observable}
	p.Subscribe(t)

	for i := 16; i <= 20; i++ {
		fmt.Println("Setting the age to", i)
		p.SetAge(i)
	}
}
