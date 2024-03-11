package main

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func NewCreature(name string, attack int, defense int) *Creature {
	return &Creature{name, attack, defense}
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)",
		c.Name, c.Attack, c.Defense)
}
