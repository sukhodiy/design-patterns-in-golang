package main

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}
