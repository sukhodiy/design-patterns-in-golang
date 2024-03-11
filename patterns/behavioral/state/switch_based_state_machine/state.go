package main

type State int

const (
	Locked State = iota
	Failed
	Unlocked
)
