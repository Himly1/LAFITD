package utils

func IsContain(values []string, value string) bool {
	for _, e := range values {
		if e == value {
			return true
		}
	}

	return false
}