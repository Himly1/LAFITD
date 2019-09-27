package util

type BadValueTypeError struct {
	CanHandledType string
}

func (BadValueTypeError BadValueTypeError) Error()string  {
	return "Can`t handle this type. Please make sure the type is : " + BadValueTypeError.CanHandledType
}

type BadConditionError struct {
	Msg string
}

func (BadConditionError BadConditionError) Error() string {
	return "Bad condition : " + BadConditionError.Msg
}

type ConditionUninitializedError struct {

}

func (ConditionUninitializedError) Error () string  {
	return "Please init the condition before do the operation"
}