package stringslice

import "gitlab.com/auk-go/core/constants"

// SafeIndexRanges get safe index range values
// If index range is out of range or not found then empty string will be added to response
//
// Reference Example :
// https://play.golang.org/p/GifVBDSqTJ2
func SafeIndexRanges(
	slice []string,
	startIndexIncluding, endIndexIncluding int,
) []string {
	lastIndex := len(slice) - 1
	requestLength := endIndexIncluding - startIndexIncluding + 1 // +1 for endingIndex or same one if zero

	if requestLength < 0 {
		return []string{}
	}

	responseSlice := MakeLen(requestLength)
	if lastIndex < constants.Zero || requestLength <= 0 {
		return responseSlice
	}

	index := -1
	for i := startIndexIncluding; i <= endIndexIncluding; i++ {
		index++

		if i > lastIndex || i < 0 {
			continue
		}

		responseSlice[index] = slice[i]
	}

	return responseSlice
}
