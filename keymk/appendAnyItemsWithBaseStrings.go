package keymk

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func appendAnyItemsWithBaseStrings(
	isSkipEmpty bool,
	mainSlice []string,
	appendingItems []interface{},
) []string {
	if len(appendingItems) == 0 {
		return mainSlice
	}

	for _, item := range appendingItems {
		if item == nil {
			continue
		}

		val := fmt.Sprintf(
			constants.SprintValueFormat,
			item)

		if isSkipEmpty && val == "" {
			continue
		}

		mainSlice = append(
			mainSlice,
			val)
	}

	return mainSlice
}
