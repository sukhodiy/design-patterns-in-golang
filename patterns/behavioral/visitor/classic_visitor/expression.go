package main

type Expression interface {
	// Double dispatch to take pointer to current expression
	Accept(ev ExpressionVisitor)
}
