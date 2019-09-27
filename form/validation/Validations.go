package validation

import (
	"pm/erros/form/validation"
	"pm/expression/core"
)

//表达式验证. 若表达式为真则验证通过， 反之则验证失败
type ExpressionValidation struct {
	expression *core.ExpressionDefinition
	errorMessages *ErrorMessages
}

func (ExpressionValidation *ExpressionValidation) Init(value interface{}) *validation.BadInitValue  {

}

func (ExpressionValidation *ExpressionValidation) IsInitialized() bool {

}


func (ExpressionValidation *ExpressionValidation) GetValidationType() string  {

}
