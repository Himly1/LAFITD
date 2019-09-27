package util

import (
	"github.com/sirupsen/logrus"
	"pm/erros/expression/core/condition/util"
	"pm/expression/core/condition"
	util2 "pm/expression/value/util"
	"pm/utils"
	. "strings"
)

func EnsureTheConditionIsOk(condition condition.Condition) *util.BadConditionError {
	return EnsureTheConditionDefinitionIsOk(condition.GetDefinition())
}

func EnsureTheConditionDefinitionIsOk(definition *condition.ConditionDefinition) *util.BadConditionError {
	badLeftValueError := util2.EnsureTheValueIsOk(definition.LeftValue)
	if badLeftValueError != nil {
		return &util.BadConditionError{Msg: "Bad leftValue. Msg : " + badLeftValueError.Error()}
	}

	badRightValueError := util2.EnsureTheValueIsOk(definition.RightValue)
	if badRightValueError != nil {
		return &util.BadConditionError{Msg: "Bad rightValue. Msg : " + badRightValueError.Error()}
	}
	
	leftValueType := util2.GetValueType(definition.LeftValue)
	rightValueType := util2.GetValueType(definition.RightValue)
	return ensureTheConditionMeetsTheStandard(leftValueType, definition.OperationType, rightValueType)
}

func ensureTheConditionMeetsTheStandard(leftValueType string, operationType string, rightValueType string) *util.BadConditionError {
	logrus.Debug("The type of the leftValue is :", leftValueType, " The operationType is: ", operationType, " The type of the rightValue is: ", rightValueType)

	legalOperations, legalOperationsIsFound := condition.GetStandard()[leftValueType]
	if !legalOperationsIsFound {
		return &util.BadConditionError{Msg: "No such standard found"}
	}

	legalRightValueTypes, legalRightValueTypesFound := legalOperations[operationType]
	if !legalRightValueTypesFound {
		allowedOperations := utils.GetKeys(legalOperations)
		return &util.BadConditionError{Msg: "Bad operationType. The allowed operations are : " + Join(allowedOperations, "")}
	}


	contains := utils.IsContain(legalRightValueTypes, rightValueType)
	if !contains {
		return &util.BadConditionError{Msg:"Bad rightValue. The allowed types of rightValue are: " + Join(legalRightValueTypes, "")}
	}

	return nil
}


func EnsureConditionInitialized(condition condition.Condition) *util.ConditionUninitializedError  {
	if !condition.IsInitialized() {
		return &util.ConditionUninitializedError{}
	}

	return nil
}