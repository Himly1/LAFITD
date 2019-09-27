package condition

type UnsupportedValueTypeError struct {
	TheType string
}

func (UnsupportedValueTypeError UnsupportedValueTypeError) Error() string  {
	return "The type of the value  does not support yet. The type is : " + UnsupportedValueTypeError.TheType
}

type UnsupportedOperationTypeError struct {
	TheType string
}

func (UnsupportedOperationTypeError UnsupportedOperationTypeError)Error () string  {
	return "The operation type of the condition  does not support yet. The type is " + UnsupportedOperationTypeError.TheType
}