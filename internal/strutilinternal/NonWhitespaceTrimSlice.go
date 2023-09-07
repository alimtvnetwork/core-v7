package strutilinternal

import "strings"

func NonWhitespaceTrimSlice(slice []string) []string {
	if slice == nil {
		return []string{}
	}

	length := len(slice)

	if length == 0 {
		return []string{}
	}

	newSlice := make(
		[]string,
		0,
		length)

	for _, s := range slice {
		trimmed := strings.TrimSpace(s)
		if len(trimmed) == 0 {
			continue
		}

		newSlice = append(newSlice, trimmed)
	}

	return newSlice
}
