package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Example: state machine for phone's states.

func main() {
	state, exitState := OffHook, OnHook
	for state != exitState {
		fmt.Println("The phone is currently", state)
		fmt.Println("Select a trigger:")

		for i := 0; i < len(rules[state]); i++ {
			tr := rules[state][i]
			fmt.Println(strconv.Itoa(i), ".", tr.Trigger)
		}

		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		i, _ := strconv.Atoi(string(input))

		tr := rules[state][i]
		state = tr.State
	}
	fmt.Println("We are done using the phone")
}
