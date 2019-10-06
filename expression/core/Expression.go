package core

import "LAFITD/expression/core/condition"

type ExpressionDefinition struct {
	Head *ExpressionNode
}

type ExpressionNode struct {
	Condition           condition.Condition
	RightSideExpression *ExpressionNode
	// unset: -1  or: 0  and: 1
	OperationType int
}

const OperationTypeUnset = -1

const OperationTypeOr = 0

const OperationTypeAnd = 1
