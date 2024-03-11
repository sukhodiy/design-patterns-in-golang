package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/sukhodiy/bookings/design-patterns-in-go/catalog/creational/prototype/copy_through_serialization/somepkg"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	//fmt.Println(b.String())

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)

	return &result
}

func testFunc() {
	a := somepkg.NewSomeStruct()
	fmt.Println(a)
	b := a.DeepCopy()
	fmt.Println(b)

	/*v := reflect.ValueOf(b).Elem()
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Type().Field(i).Name, v.Field(i))
	}*/

}

func main() {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt"},
	}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)

	testFunc()
}
