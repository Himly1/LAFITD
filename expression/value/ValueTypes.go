package value

import "pm/erros/expression/core/condition"

var valueTypes = map[string]string {
	SubjectPropertyType:SubjectPropertyType,
	BoolType:BoolType,
	StringType:StringType,
	IntType: IntType,
	FloatType:FloatType,
	CollectionStringType:CollectionStringType,
	CollectionIntType:CollectionIntType,
	CollectionFloatType:CollectionFloatType,
}

const SubjectPropertyType = "subjectProperty"

const BoolType = "bool"

const StringType = "string"

const IntType = "int"

const FloatType = "float"

const CollectionStringType = "collectionString"

const CollectionIntType = "collectionInt"

const CollectionFloatType = "CollectionFloat"

func EnsureTypeIsExists(conditionType string) *condition.UnsupportedValueTypeError {
	_, found := valueTypes[conditionType]
	if !found {
		return &condition.UnsupportedValueTypeError{TheType: conditionType}
	}

	return nil
}
