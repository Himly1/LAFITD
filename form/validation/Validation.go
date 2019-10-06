package validation

import (
	"lafitd/erros/form/validation"
	international_message "lafitd/international-message"
)

type Validation interface {
	Init(value interface{}) *validation.BadInitValue
	IsInitialized() bool
	GetValidationType() string
}

//异常信息。 用户定义验证不通过时的异常信息.
type ErrorMessages struct {
	messages []*international_message.Message
}




