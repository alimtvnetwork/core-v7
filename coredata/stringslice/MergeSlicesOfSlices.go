package stringslice

import "gitlab.com/auk-go/core/constants"

// MergeSlicesOfSlices Don't include nil or length 0 slices
func MergeSlicesOfSlices(slicesOfSlice ...[]string) []string {
	if len(slicesOfSlice) == 0 {
		return []string{}
	}

	sliceLength := len(slicesOfSlice)

	if sliceLength == constants.Zero {
		return []string{}
	}

	countOfAll := AllElemLengthSlices(slicesOfSlice...)

	if countOfAll == constants.Zero {
		return []string{}
	}

	newSlice := make(
		[]string,
		constants.Zero,
		countOfAll)

	for _, slice := range slicesOfSlice {
		if len(slice) == 0 {
			continue
		}

		newSlice = append(newSlice, slice...)
	}

	return newSlice
}
