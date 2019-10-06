package validation

import (
	"github.com/sirupsen/logrus"
	"LAFITD/erros/form/validation"
	"LAFITD/expression/core"
	"sync"
)

//表达式验证. 若表达式为真则验证通过， 反之则验证失败
type ExpressionValidation struct {
	Expression *core.ExpressionDefinition
	ErrorMessages *ErrorMessages
	initialized bool
	lock sync.RWMutex
}

func (ExpressionValidation *ExpressionValidation) Init(value interface{}) *validation.BadInitValue  {
	logrus.WithFields(logrus.Fields{
		"value":value,
	}).Debug("Received value is: ")

	ExpressionValidation.lock.Lock()
	defer ExpressionValidation.lock.Unlock()

	if ExpressionValidation.initialized {
		return nil
	}

	realValue, ok := value.(*core.ExpressionDefinition)
	if ok {
		ExpressionValidation.Expression = realValue
		return nil
	}else {
		logrus.Error("The value cannot assert to ExpressionDefinition")
		return &validation.BadInitValue{}
	}
}

func (ExpressionValidation *ExpressionValidation) IsInitialized() bool {
	return ExpressionValidation.initialized
}


func (ExpressionValidation *ExpressionValidation) GetValidationType() string  {
	return ALL_EXPRESSION_VALIDATION
}
