package main

import "fmt"

// Example: money transfer command for bank account with Undo function.
// It contains 2 commands inside - Withdraw and Deposit, and also it has
// Succeeded status of command.

// Composite command consists several diff commands

func main() {
	from := BankAccount{100}
	to := BankAccount{0}
	mtc := NewMoneyTransferCommand(&from, &to, 25)
	mtc.Call()
	fmt.Println(from, to)

	mtc.Undo()
	fmt.Println(from, to)
}
