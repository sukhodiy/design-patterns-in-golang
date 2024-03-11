package main

type ExpressionEvaluator struct {
	result float64
}

func (ee *ExpressionEvaluator) VisitDoubleExpression(e *DoubleExpression) {
	ee.result = e.value
}

func (ee *ExpressionEvaluator) VisitAdditionExpression(e *AdditionExpression) {
	e.left.Accept(ee)
	x := ee.result
	e.right.Accept(ee)
	x += ee.result
	ee.result = x
}
