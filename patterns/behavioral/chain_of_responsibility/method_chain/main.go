package main

import (
	"fmt"
)

// Example: creature's modifier for game.

func main() {
	goblin := NewCreature("Goblin", 1, 1)

	fmt.Println(goblin.String())

	root := NewCreatureModifier(goblin)

	//root.Add(NewNoBonusesModifier(goblin))

	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewIncreaseDefenseModifier(goblin))
	root.Add(NewDoubleAttackModifier(goblin))

	root.Handle()

	// root := NewDoubleAttackModifier(goblin)
	// root.Add(NewIncreaseDefenseModifier(goblin))
	// root.Handle()

	fmt.Println(goblin.String())
}
