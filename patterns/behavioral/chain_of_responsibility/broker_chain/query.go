package main

type Argument int

const (
	Attack Argument = iota
	Defense
)

// CQS
type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}
