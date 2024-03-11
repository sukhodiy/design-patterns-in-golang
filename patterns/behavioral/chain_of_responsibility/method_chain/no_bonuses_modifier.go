package main

type NoBonusesModifier struct {
	CreatureModifier
}

func NewNoBonusesModifier(c *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{
		CreatureModifier{creature: c},
	}
}

func (n *NoBonusesModifier) Handle() {
	// empty. It breaks chain.
}
