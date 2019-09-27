package international_message

type CountryNotFoundError struct {
	Code string
}

func (CountryNotFoundError CountryNotFoundError) Error () string  {
	return "No such country exists. The code is " + CountryNotFoundError.Code
}

