package util

import (
	"pm/erros/expression/core/condition/util"
	"pm/expression/value"
	Values  "pm/expression/value"
)

func GetValueAsBoolValue(value Values.Value) (*util.BadValueTypeError, *Values.BoolValue) {
	err := ensureValueCanBeHandled(value.GetType(), Values.BoolType)
	if err != nil {
		return err, nil
	}else {
		return nil,value.(*Values.BoolValue)
	}
}

func GetValueAsSubjectProperty(value Values.Value) (*util.BadValueTypeError, *Values.SubjectPropertyValue) {
	err := ensureValueCanBeHandled(value.GetType(), Values.SubjectPropertyType)
	if err != nil {
		return err, nil
	}else {
		return nil, value.(*Values.SubjectPropertyValue)
	}
}

func GetValueAsIntValue(value value.Value) (*util.BadValueTypeError, *Values.IntValue)  {
	err := ensureValueCanBeHandled(value.GetType(), Values.IntType)
	if err != nil {
		return err, nil
	}else {
		return nil, value.(*Values.IntValue)
	}
}

func GetValueAsStringValue(value value.Value) (*util.BadValueTypeError, *Values.StringValue)  {
	err := ensureValueCanBeHandled(value.GetType(), Values.StringType)
	if err != nil {
		return err, nil
	}else {
		return nil, value.(*Values.StringValue)
	}
}

func GetValueAsFloatValue(value value.Value) (*util.BadValueTypeError, *Values.FloatValue) {
	err := ensureValueCanBeHandled(value.GetType(), Values.FloatType)
	if err != nil {
		return err, nil;
	}else {
		return nil, value.(*Values.FloatValue)
	}
}

func GetValueAsCollectionIntValue(value value.Value) (*util.BadValueTypeError, *Values.CollectionIntValue) {
	err := ensureValueCanBeHandled(value.GetType(), Values.CollectionIntType)
	if err != nil {
		return err, nil
	}else {
		return nil, value.(*Values.CollectionIntValue)
	}
}

func GetValueAsCollectionStringValue(value value.Value) (*util.BadValueTypeError, *Values.CollectionStringValue) {
	err := ensureValueCanBeHandled(value.GetType(), Values.CollectionStringType)
	if err != nil {
		return err ,nil
	}else {
		return nil, value.(*Values.CollectionStringValue)
	}
}

func GetValueAsCollectionFloatValue(value value.Value) (*util.BadValueTypeError, *Values.CollectionFloatValue) {
	err := ensureValueCanBeHandled(value.GetType(), Values.CollectionFloatType)
	if err != nil {
		return err, nil
	}else {
		return nil, value.(*Values.CollectionFloatValue)
	}
}

func ensureValueCanBeHandled(valueType string, canHandledType string) *util.BadValueTypeError {
	if valueType != canHandledType {
		return &util.BadValueTypeError{CanHandledType: canHandledType}
	}

	return nil
}