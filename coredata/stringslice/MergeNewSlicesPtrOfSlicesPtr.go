package stringslice

import "gitlab.com/evatix-go/core/constants"

// MergeNewSlicesPtrOfSlicesPtr Don't include nil or length 0 slices
func MergeNewSlicesPtrOfSlicesPtr(slices *[]*[]string) *[]string {
	if slices == nil {
		return &[]string{}
	}

	sliceLength := len(*slices)

	if sliceLength == constants.Zero {
		return &[]string{}
	}

	countOfAll := AllElemLengthSlicesPtr(slices)

	if countOfAll == constants.Zero {
		return &[]string{}
	}

	newSlice := make(
		[]string,
		constants.Zero,
		countOfAll)

	for _, slice := range *slices {
		if len(*slice) == constants.Zero {
			continue
		}

		newSlice = append(newSlice, *slice...)
	}

	return &newSlice
}
