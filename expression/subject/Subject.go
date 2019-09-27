package subject

import "pm/erros/expression/subject"

//表达的主体
type Subject interface {
	GetSubjectInfo() *BasicSubjectInfo

	//新增属性
	AddProperty(property *Property) *subject.DuplicatedError
}

type BasicSubjectInfo struct {
	Code string
	Name string
	Description string
	Properties []*Property
}


//主体的属性
type Property struct {
	Code string

	Name string

	//属性类型.
	ValueType string

	//当前的属性是否是另一个主体
	IsSubject bool

	//当isSubject为true时该地址变量指向的是一个新的主体.
	Subject *Subject
}


type DynamicSubject struct {
	basicInfo *BasicSubjectInfo
}

func (DynamicSubject *DynamicSubject) GetSubjectInfo() *BasicSubjectInfo  {
	return DynamicSubject.basicInfo
}

func (DynamicSubject *DynamicSubject) AddProperty(property *Property) *subject.DuplicatedError {
	properties := DynamicSubject.basicInfo.Properties
	duplicatedError := ensureNoDuplicatedProperty(property, properties)
	if duplicatedError != nil {
		return duplicatedError
	}

	DynamicSubject.basicInfo.Properties  = append(properties, property)
	return nil
}

func ensureNoDuplicatedProperty(property *Property, properties []*Property) *subject.DuplicatedError {
	for _, e := range properties  {
		if e.Code == property.Code {
			return &subject.DuplicatedError{Code:e.Code}
		}
	}

	return nil
}
