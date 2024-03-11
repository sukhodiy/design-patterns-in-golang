package main

type Observer interface {
	Notify(data interface{})
}
