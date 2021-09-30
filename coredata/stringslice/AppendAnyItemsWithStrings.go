package stringslice

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func AppendAnyItemsWithStrings(
	isClone,
	isSkipOnEmpty bool,
	mainSlice []string,
	appendingItems ...interface{},
) []string {
	slice := CloneIf(
		isClone,
		len(appendingItems)+constants.Capacity2,
		mainSlice)

	if len(appendingItems) == 0 {
		return slice
	}

	for _, item := range appendingItems {
		if item == nil {
			continue
		}

		val := fmt.Sprintf(
			constants.SprintValueFormat,
			item)

		if isSkipOnEmpty && val == "" {
			continue
		}

		slice = append(
			slice,
			val)
	}

	return slice
}
