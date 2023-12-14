package stringslice

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// ExpandBySplits
// Take each slice item, split and add to the new slice array and returns it.
func ExpandBySplits(
	slice []string,
	splitters ...string,
) []string {
	length := len(slice)
	if length == 0 {
		return []string{}
	}

	splitExpandFunc := func(line string) []string {
		if len(splitters) == 0 {
			return []string{}
		}

		newExpandedSlice := make([]string, 0, constants.Capacity8)

		for _, splitter := range splitters {
			lines := strings.Split(line, splitter)
			newExpandedSlice = append(newExpandedSlice, lines...)
		}

		return newExpandedSlice
	}

	expandedSlicesOfSlice := ExpandByFunc(slice, splitExpandFunc)

	return expandedSlicesOfSlice
}
