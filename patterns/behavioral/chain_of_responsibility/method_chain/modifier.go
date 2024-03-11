package main

type Modifier interface {
	Add(m Modifier)
	Handle()
}
