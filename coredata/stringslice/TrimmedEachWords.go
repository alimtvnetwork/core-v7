package stringslice

import (
	"strings"
)

// TrimmedEachWords
//
// Create a final slice by trimming each words in side each line.
// No spaces remained.
func TrimmedEachWords(slice []string) []string {
	if slice == nil {
		return []string{}
	}

	length := len(slice)

	if length == 0 {
		return []string{}
	}

	newSlice := MakeDefault(length)

	for _, s := range slice {
		trimmed := strings.TrimSpace(s)
		if len(trimmed) == 0 {
			continue
		}

		newSlice = append(newSlice, trimmed)
	}

	return newSlice
}
