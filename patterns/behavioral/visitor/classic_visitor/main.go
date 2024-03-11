package main

import (
	"fmt"
)

// Example: print and evaluate math expression.

func main() {
	// 1 + (2 + 3) = 6
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}

	ep := NewExpressionPrinter()
	e.Accept(ep)
	fmt.Println(ep.String())

	ee := &ExpressionEvaluator{}
	e.Accept(ee)
	fmt.Printf("%s = %g\n", ep, ee.result)
}
