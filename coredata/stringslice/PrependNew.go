package stringslice

import "gitlab.com/auk-go/core/constants"

func PrependNew(
	secondSlice []string,
	prependingItems ...string,
) *[]string {
	sliceLength := len(secondSlice)
	additionalItemsLength := len(prependingItems)

	newSlice := make(
		[]string,
		constants.Zero,
		sliceLength+additionalItemsLength)

	if additionalItemsLength > 0 {
		newSlice = append(newSlice, prependingItems...)
	}

	if sliceLength > 0 {
		newSlice = append(newSlice, secondSlice...)
	}

	return &newSlice
}
