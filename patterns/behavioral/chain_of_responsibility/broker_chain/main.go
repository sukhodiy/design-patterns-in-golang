package main

import (
	"fmt"
	"sync"
)

// Using patterns:
// CoR, Mediator, Observer, CQS

// This approach helps us calculate Creature properties during methods calling.

func main() {
	game := &Game{sync.Map{}}

	goblin := NewCreature(game, "Strong Goblin", 2, 2)
	fmt.Println(goblin.String())

	{
		m := NewDoubleAttackModifier(game, goblin)
		fmt.Println(goblin.String())
		m.Close()
	}

	fmt.Println(goblin.String())
}
