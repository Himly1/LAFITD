package core

import (
	"LAFITD/erros/expression/core"
	"LAFITD/expression/core/condition"
	"sync"
)


type ExpressionBuilder interface {
	SetFirstCondition(condition condition.Condition)
	Or(condition condition.Condition) *core.FirstConditionUnsetError
	And(Condition condition.Condition) *core.FirstConditionUnsetError
	GetExpressionDefinition() (*core.FirstConditionUnsetError, *ExpressionDefinition)
}

type DefaultExpressionBuilder struct {
	expressionDefinition *ExpressionDefinition
	lock *sync.RWMutex
}

func (Builder  *DefaultExpressionBuilder)SetFirstCondition(condition condition.Condition)  {
	Builder.lock.Lock()
	defer Builder.lock.Unlock()

	if isFirstConditionIsNil(Builder) {
		head := createExpressionNode(condition, OperationTypeUnset, nil)
		Builder.expressionDefinition.Head = head
	}
}

func (Builder *DefaultExpressionBuilder)Or(condition condition.Condition) *core.FirstConditionUnsetError  {
	Builder.lock.Lock()
	defer Builder.lock.Unlock()

	if isFirstConditionIsNil(Builder) {
		return &core.FirstConditionUnsetError{}
	}

	newNode := createExpressionNode(condition, OperationTypeUnset, nil)
	latestNode := getLatestExpressionNode(Builder.expressionDefinition.Head)
	latestNode.OperationType = OperationTypeOr
	latestNode.RightSideExpression = newNode

	return nil
}

func createExpressionNode(condition condition.Condition, operationType int, rightSideCondition *ExpressionNode) *ExpressionNode {
	return &ExpressionNode{Condition:condition, RightSideExpression:rightSideCondition, OperationType:operationType}
}

func (Builder *DefaultExpressionBuilder)And(condition condition.Condition) *core.FirstConditionUnsetError  {
	Builder.lock.Lock()
	defer Builder.lock.Unlock()

	if isFirstConditionIsNil(Builder) {
		return &core.FirstConditionUnsetError{}
	}

	newNode := createExpressionNode(condition, OperationTypeUnset,nil)
	latestNode := getLatestExpressionNode(Builder.expressionDefinition.Head)
	latestNode.OperationType = OperationTypeAnd
	latestNode.RightSideExpression = newNode

	return nil
}

func getLatestExpressionNode(node *ExpressionNode) *ExpressionNode {
	if node.RightSideExpression == nil {
		return node
	}else {
		return getLatestExpressionNode(node.RightSideExpression)
	}
}


func (Builder *DefaultExpressionBuilder)GetExpressionDefinition() (*core.FirstConditionUnsetError, *ExpressionDefinition)  {
	Builder.lock.RLock()
	defer Builder.lock.RUnlock()

	if isFirstConditionIsNil(Builder) {
		return &core.FirstConditionUnsetError{}, nil
	}

	return nil, Builder.expressionDefinition
}

func isFirstConditionIsNil(builder *DefaultExpressionBuilder) bool {
	return builder.expressionDefinition.Head == nil
}
