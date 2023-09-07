package stringslice

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/defaultcapacity"
)

// SafeIndexesDefaultWithDetail Only indexes which are present values will be included.
//
// Warning : Not found indexes will not be included in the values.
func SafeIndexesDefaultWithDetail(
	slice []string,
	indexes ...int,
) *IndexValuesDetail {
	length := len(slice)

	if length == 0 {
		return InvalidIndexValuesDetail()
	}

	values := make(
		[]string,
		constants.Zero,
		length)
	missingIndexes := make(
		[]int,
		constants.Zero,
		defaultcapacity.OfSearch(length))

	inputIndex := 0
	lastIndex := length - 1
	for _, index := range indexes {
		if index > lastIndex || index < 0 {
			missingIndexes = append(
				missingIndexes,
				index)
			// don't include
			continue
		}

		values = append(values, slice[index])
		inputIndex++
	}

	return &IndexValuesDetail{
		Values:         values,
		MissingIndexes: missingIndexes,
		IsAnyMissing:   len(missingIndexes) > 0,
		IsValid:        true,
	}
}
