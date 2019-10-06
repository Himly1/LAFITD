package international_message

import international_message "lafitd/erros/international-message"

var countries = map[string]*Country {
	CHINA: &Country{code:CHINA},
	UNITE_STATE: &Country{code:CHINA},
}

type Country struct {
	code string
}

func GetCountry(code string) (*international_message.CountryNotFoundError, *Country) {
	country, found := countries[code]
	if !found {
		return &international_message.CountryNotFoundError{Code:code}, nil
	}else {
		return nil, country
	}
}

const CHINA = "CN"

const UNITE_STATE = "US"
