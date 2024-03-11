package main

type Integer struct {
	value int
}

func NewInteger(value int) *Integer {
	return &Integer{value: value}
}

func (i Integer) Value() int {
	return i.value
}
