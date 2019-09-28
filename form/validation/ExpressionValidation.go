package validation

import (
	"pm/erros/form/validation"
	"pm/expression/core"
)

//表达式验证. 若表达式为真则验证通过， 反之则验证失败
type ExpressionValidation struct {
	expression *core.ExpressionDefinition
	errorMessages *ErrorMessages
	initialized bool
}

func (ExpressionValidation *ExpressionValidation) Init(value interface{}) *validation.BadInitValue  {
	if ExpressionValidation.initialized {
		return nil
	}

	realValue, ok := value.(*core.ExpressionDefinition)
	if ok {
		ExpressionValidation.expression = realValue
		return nil
	}else {
		return &validation.BadInitValue{}
	}
}

func (ExpressionValidation *ExpressionValidation) IsInitialized() bool {
	return ExpressionValidation.initialized
}


func (ExpressionValidation *ExpressionValidation) GetValidationType() string  {
	return EXPRESSION_VALIDATION
}
