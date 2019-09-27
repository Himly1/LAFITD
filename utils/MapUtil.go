package utils

func GetKeys(aMap map[string][]string) []string {
	keys := make([]string, 0, len(aMap))
	for k := range aMap {
		keys = append(keys, k)
	}

	return keys
}