package main

type Command interface {
	Call()
	Undo()
	Succeeded() bool
	SetSucceeded(value bool)
}
