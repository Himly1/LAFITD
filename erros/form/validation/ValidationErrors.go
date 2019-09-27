package validation

type BadInitValue struct {
	jsonDefinition string
}

func (BadInitValue BadInitValue) Error() string  {
	return "Bad value for init the validation. Here is the example: " + BadInitValue.jsonDefinition
}

