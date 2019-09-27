package condition

import "pm/erros/expression/core/condition"

var operationTypes = map[string]string {
	Equal:Equal,
	GreaterThan:GreaterThan,
	LesserThan:LesserThan,
	Contain:Contain,
}

const Equal = "equal"

const GreaterThan = "greaterThan"

const LesserThan = "lesserThan"

const Contain = "contain"

func EnsureTheTypeIsExists(operationType string) *condition.UnsupportedOperationTypeError  {
	_, found := operationTypes[operationType]
	if !found {
		return &condition.UnsupportedOperationTypeError{}
	}

	return nil
}