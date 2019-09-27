package util


type BadValueError struct {
	Msg string
}

func (BadValueError BadValueError) Error() string {
	return "The value is`t correct. Message: " + BadValueError.Msg
}