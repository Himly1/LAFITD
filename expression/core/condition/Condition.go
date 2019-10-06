package condition

import (
	util2 "lafitd/erros/expression/core/condition/util"
	"lafitd/expression/core/condition/util"
	"lafitd/expression/value"
)


type Condition interface {
	Init(leftValue value.Value, operationType string, RightValue value.Value) *util2.BadConditionError
	IsInitialized() bool
	GetDefinition() *ConditionDefinition
}


type ConditionDefinition struct {
	LeftValue     value.Value
	OperationType string
	RightValue    value.Value
}

type DefaultCondition struct {
	initialized   bool
	leftValue     value.Value
	operationType string
	rightValue    value.Value
}

func (condition *DefaultCondition) Init(leftValue value.Value, operationType string, rightValue value.Value) *util2.BadConditionError{
	err := util.EnsureTheConditionIsOk(condition)
	if err != nil {
		return err
	}

	if !condition.initialized {
		condition.operationType = operationType
		condition.leftValue =  leftValue
		condition.rightValue = rightValue
		condition.initialized = true
	}

	return nil
}

func (condition *DefaultCondition) IsInitialized() bool  {
	return condition.initialized
}

func (condition *DefaultCondition) GetDefinition() *ConditionDefinition  {
	return &ConditionDefinition{LeftValue:condition.leftValue, OperationType:condition.operationType, RightValue:condition.rightValue}
}

