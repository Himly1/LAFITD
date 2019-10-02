package form

import "pm/form/validation"

type Definition struct {
	items []*Definition
	options []*Option
	basicInfo *BasicDefinition
	frontEndValidations []validation.Validation
	backendValidations []validation.Validation
}