package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited", amount,
		"\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount < overdraftLimit {
		return false
	}

	b.balance -= amount
	fmt.Println("Withdrew", amount,
		"\b, balance is now", b.balance)

	return true
}
