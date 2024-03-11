package main

import "fmt"

type IncreasedDefenseModifier struct {
	CreatureModifier
}

func NewIncreaseDefenseModifier(c *Creature) *IncreasedDefenseModifier {
	return &IncreasedDefenseModifier{CreatureModifier{creature: c}}
}

func (i *IncreasedDefenseModifier) Handle() {
	if i.creature.Attack <= 2 {
		fmt.Println("Increasing", i.creature.Name, "\b's defense")
		i.creature.Defense++
	}
	i.CreatureModifier.Handle()
}
