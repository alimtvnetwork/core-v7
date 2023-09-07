package strutilinternal

import "gitlab.com/auk-go/core/constants"

func NonEmptySlice(slice []string) []string {
	length := len(slice)

	if length == 0 {
		return []string{}
	}

	newSlice := make(
		[]string,
		0,
		length)

	for _, s := range slice {
		if s != constants.EmptyString {
			newSlice = append(newSlice, s)
		}
	}

	return newSlice
}
