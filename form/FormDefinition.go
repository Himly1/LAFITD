package form

import "pm/form/validation"

type Definition struct {
	items []*Definition
	options []*Option
	frontEndValidations []validation.Validation
	backendValidations []validation.Validation
}