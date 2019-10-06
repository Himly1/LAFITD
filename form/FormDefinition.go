package form

import "lafitd/form/validation"

type Definition struct {
	items []*Definition
	options []*Option
	basicInfo *BasicDefinition
	frontEndValidations []validation.Validation
	backendValidations []validation.Validation
}