package main

import (
	"fmt"
	"strings"
)

// String Builder is visitor here
type Expression interface {
	Print(sb *strings.Builder) // but this is violating Single Responsibility Principle
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", d.value))
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteRune('(')
	a.left.Print(sb)
	sb.WriteRune('+')
	a.right.Print(sb)
	sb.WriteRune(')')
}

func main() {
	// 1 + (2 + 3)
	e := AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	sb := strings.Builder{} // String Builder is visitor here
	e.Print(&sb)
	fmt.Println(sb.String())
}
