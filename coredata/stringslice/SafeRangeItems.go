package stringslice

import "gitlab.com/evatix-go/core/constants"

func SafeRangeItems(
	slice []string,
	start, end int,
) []string {
	if slice == nil {
		return []string{}
	}

	length := len(slice)
	if length == 0 {
		return []string{}
	}

	lastIndex := length - 1
	if start > lastIndex {
		return []string{}
	}

	if end > lastIndex || end == constants.InvalidValue {
		end = lastIndex
	}

	if start <= lastIndex && start != constants.InvalidValue {
		return slice[start:end]
	}

	return slice[:end]
}
