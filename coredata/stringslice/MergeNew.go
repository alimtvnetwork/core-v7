package stringslice

import (
	"gitlab.com/evatix-go/core/constants"
)

func MergeNew(firstSlice []string, additionalItems ...string) *[]string {
	sliceLength := len(firstSlice)
	additionalItemsLength := len(additionalItems)

	newSlice := make(
		[]string,
		constants.Zero,
		sliceLength+additionalItemsLength)

	if sliceLength > 0 {
		newSlice = append(newSlice, firstSlice...)
	}

	if additionalItemsLength > 0 {
		newSlice = append(newSlice, additionalItems...)
	}

	return &newSlice
}
