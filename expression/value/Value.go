package value

import (
	"LAFITD/expression/subject"
)

type Value interface {
	GetType() string
}

type SubjectPropertyValue struct {
	Subject subject.Subject
	Property *subject.Property
}

func (SubjectPropertyValue *SubjectPropertyValue)GetType() string {
	return SubjectPropertyType
}

type BoolValue struct {
	Value bool
}

func (BoolValue *BoolValue) GetType () string {
	return BoolType
}

type StringValue struct {
	Value string
}

func (StringValue *StringValue) GetType () string {
	return StringType
}

type IntValue struct {
	Value int64
}

func (IntValue *IntValue) GetType () string {
	return IntType
}

type FloatValue struct {
	Value float64
}

func (FloatValue *FloatValue) GetType () string {
	return FloatType
}

type CollectionStringValue struct {
	Values []string
}

func (CollectionStringValue *CollectionStringValue) GetType () string {
	return CollectionStringType
}

type CollectionIntValue struct {
	Values []int64
}

func (CollectionIntValue *CollectionIntValue) GetType () string {
	return CollectionIntType
}

type CollectionFloatValue struct {
	Values []float64
}

func (CollectionFloatValue *CollectionFloatValue) GetType () string {
	return CollectionFloatType
}
