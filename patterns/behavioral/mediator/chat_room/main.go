package main

import (
	"fmt"
)

// Example: chat room.

// Mediator. Person does not know about other persons (no pointers).
type ChatRoom struct {
	people []*Person
}

func (c *ChatRoom) Broadcast(src, msg string) {
	for _, p := range c.people {
		if p.Name != src {
			p.Receive(src, msg)
		}
	}
}

func (c *ChatRoom) Message(src, dst, msg string) {
	for _, p := range c.people {
		if p.Name == dst {
			p.Receive(src, msg)
		}
	}
}

func (c *ChatRoom) Join(p *Person) {
	joinMsg := p.Name + " joins the chat"
	c.Broadcast("Room", joinMsg)

	p.Room = c
	c.people = append(c.people, p)
}

type Person struct {
	Name    string
	Room    *ChatRoom
	chatLog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Receive(sender, msg string) {
	s := fmt.Sprintf("%s: %s\n", sender, msg)
	fmt.Printf("[%s's chat session]: %s", p.Name, s)
	p.chatLog = append(p.chatLog, s)
}

func (p *Person) Say(msg string) {
	p.Room.Broadcast(p.Name, msg)
}

func (p *Person) PrivateMessage(whom, msg string) {
	p.Room.Message(p.Name, whom, msg)
}

func main() {
	room := ChatRoom{}

	john := NewPerson("John")
	jane := NewPerson("Jane")

	room.Join(john)
	room.Join(jane)

	john.Say("hi room")
	jane.Say("oh, hey john")

	simon := NewPerson("Simon")
	room.Join(simon)
	simon.Say("hi everyone!")

	jane.PrivateMessage("Simon", "glad you could join us!")
}
