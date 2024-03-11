package main

import "fmt"

type TrafficManagement struct {
	o Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
	pc, ok := data.(PropertyChange)
	if !ok {
		return
	}
	if pc.Value.(int) >= 18 {
		fmt.Println("Congrats, you can drive now!")
		t.o.Unsubscribe(t)
	}
}
