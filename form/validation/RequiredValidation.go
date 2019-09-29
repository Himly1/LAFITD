package validation

import (
	"github.com/sirupsen/logrus"
	"pm/erros/form/validation"
)

type RequiredValidation struct {

}

func (RequiredValidation *RequiredValidation) Init(value interface{}) *validation.BadInitValue  {

}

func (RequiredValidation *RequiredValidation) IsInitialized() bool {

}


func (RequiredValidation *RequiredValidation) GetValidationType() string  {
	return EXPRESSION_VALIDATION
}