package validation

import (
	"pm/erros/form/validation"
)

type NotNullValidation struct {

}

func (RequiredValidation *NotNullValidation) Init(value interface{}) *validation.BadInitValue  {
	return nil
}

func (RequiredValidation *NotNullValidation) IsInitialized() bool {
	return true
}


func (RequiredValidation *NotNullValidation) GetValidationType() string  {
	return ALL_NOTNUILL_VALIDATION
}