package stringslice

import "gitlab.com/evatix-go/core/constants"

func AppendStringsWithAnyItems(
	isClone,
	isSkipOnEmpty bool,
	mainSlice []interface{},
	appendingItems ...string,
) []interface{} {
	slice := AnyItemsCloneIf(
		isClone,
		len(appendingItems)+constants.Capacity2,
		mainSlice)

	if len(appendingItems) == 0 {
		return slice
	}

	for _, item := range appendingItems {
		if isSkipOnEmpty && item == "" {
			continue
		}

		slice = append(
			slice,
			item)
	}

	return slice
}
