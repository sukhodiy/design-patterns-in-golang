package main

import (
	"fmt"
)

type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinningPlayer() int
}

// Template method
func PlayGame(g Game) {
	g.Start()

	for !g.HaveWinner() {
		g.TakeTurn()
	}

	fmt.Printf("Player %d wins.\n", g.WinningPlayer())
}

type chess struct {
	turn, maxTurns, currentPlayer int
}

func NewGameOfChess() Game {
	return &chess{1, 10, 0}
}

func (c *chess) Start() {
	fmt.Println("Starting a new game of chess.")
}

func (c *chess) TakeTurn() {
	c.turn++
	fmt.Printf("Turn %d taken by playes %d\n",
		c.turn, c.currentPlayer)
	c.currentPlayer = 1 - c.currentPlayer

}

func (c *chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

func (c *chess) WinningPlayer() int {
	return c.currentPlayer
}

func main() {
	chess := NewGameOfChess()
	PlayGame(chess)
}
