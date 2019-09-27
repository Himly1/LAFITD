package form

import "pm/form/validation"

type Definition struct {
	items []*Definition
	options []*Option
	frontEndValidations []validation.BackendValidation
	backendValidations []validation.FrontEndValidation
}