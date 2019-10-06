package user_task_owner

type BadUserOwnerError struct {
	Msg string
}

func (BadUserOwnerError *BadUserOwnerError) Error() string {
	return BadUserOwnerError.Msg
}