package main

import "strconv"

func Parse(tokens []Token) Element {
	result := BinaryOperation{}
	haveLhs := false
	for i := 0; i < len(tokens); i++ {
		token := &tokens[i]
		switch token.Type {
		case Int:
			n, _ := strconv.Atoi(token.Text)
			integer := Integer{n}
			if !haveLhs {
				result.Left = &integer
				haveLhs = true
			} else {
				result.Right = &integer
			}
		case Plus:
			result.Type = Addition
		case Minus:
			result.Type = Subtraction
		case Lparen:
			j := i
			for ; j < len(tokens); j++ {
				if tokens[j].Type == Rparen {
					break
				}
			}
			// now j points to closing bracket, so
			// process subexpression without opening
			var subexp []Token
			for k := i + 1; k < j; k++ {
				subexp = append(subexp, tokens[k])
			}
			element := Parse(subexp)
			if !haveLhs {
				result.Left = element
				haveLhs = true
			} else {
				result.Right = element
			}
			i = j
		}
	}
	return &result
}
