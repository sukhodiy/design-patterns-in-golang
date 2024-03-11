package main

import (
	"fmt"
)

func main() {
	input := "(13+4)-(12+1)"
	tokens := Lex(input)
	fmt.Println(tokens)
}
