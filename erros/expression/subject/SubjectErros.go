package subject

type NotFoundError struct {
	Code string
}

func (SubjectNotFoundError *NotFoundError) Error() string  {
	return "No such subject found. The code is " + SubjectNotFoundError.Code
}

type DuplicatedError struct {
	Code string
}

func (SubjectDuplicatedError DuplicatedError) Error() string {
	return "Found duplicated subject. The code is " + SubjectDuplicatedError.Code
}

