package main

import "fmt"

type Creature struct {
	game            *Game
	Name            string
	attack, defense int // only initial values
}

func NewCreature(game *Game, name string, attack int,
	defense int) *Creature {
	return &Creature{game, name, attack, defense}
}

func (c *Creature) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) Defense() int {
	q := Query{c.Name, Defense, c.defense}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack(),
		c.Defense())
}
