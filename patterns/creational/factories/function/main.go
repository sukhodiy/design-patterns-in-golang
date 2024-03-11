package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	if age < 16 {
		// ...
	}
	return &Person{name, age, 2} // 2 - default parameters
}

func main() {
	p := NewPerson("John", 33)
	//p.EyeCount = 1
	fmt.Println(p)
}
