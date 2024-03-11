package main

import (
	"container/list"
	"fmt"
)

// Observable, Observer

type Observer interface {
	Notify(data interface{})
}

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type ProperyChange struct {
	Name  string // "Age"
	Value interface{}
}

type Person struct {
	Observable
	age int // Age() SetAge()
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}

	oldCanVote := p.CanVote()

	p.age = age
	p.Fire(ProperyChange{"Age", p.age})

	if oldCanVote != p.CanVote() {
		p.Fire(ProperyChange{"CanVote", p.CanVote()})
	}
}

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

func (p *Person) CanVote() bool {
	return p.age >= 18
}

type ElectoralRoll struct {
}

func (e *ElectoralRoll) Notify(data interface{}) {
	if pc, ok := data.(ProperyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congratulations, you can vote!")
		}
	}
}

func main() {
	p := NewPerson(0)
	er := &ElectoralRoll{}
	p.Subscribe(er)

	for i := 10; i < 20; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}

}
