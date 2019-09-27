package util

import (
	"pm/erros/expression/core/value/util"
	util2 "pm/expression/core/condition/util"
	"pm/expression/subject"
	"pm/expression/value"
)

var checkers = map[string]func(value value.Value) *util.BadValueError {
	value.SubjectPropertyType: ensureTheSubjectPropertyValueIsOk,
}

func EnsureTheValueIsOk(value value.Value) *util.BadValueError {
	checker, found := checkers[value.GetType()]
	if found {
		return checker(value)
	}

	return nil
}

func ensureTheSubjectPropertyValueIsOk(value value.Value) *util.BadValueError {
	_, realValue := util2.GetValueAsSubjectProperty(value)
	return ensureThePropertyIsTheOneOfThePropertiesOfTheSubject(realValue.Property, realValue.Subject)
}

func ensureThePropertyIsTheOneOfThePropertiesOfTheSubject(property *subject.Property, aSubject subject.Subject) *util.BadValueError  {
	flag := false

	for _, e := range aSubject.GetSubjectInfo().Properties {
		codeEqual := e.Code == property.Code
		nameEqual := e.Name == property.Name

		if codeEqual && nameEqual {
			flag = true
			break
		}
	}

	if !flag {
		return &util.BadValueError{Msg: "The property must be the one of the properties of the subject"}
	}else {
		return nil
	}
}

func GetValueType(aValue value.Value) string {
	if aValue.GetType() == value.SubjectPropertyType {
		_, realValue := util2.GetValueAsSubjectProperty(aValue)
		return realValue.Property.ValueType
	}else {
		return aValue.GetType()
	}
}